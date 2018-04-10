package volumes

import (
	"context"
	"fmt"
	"time"

	"github.com/signalfx/golib/datapoint"
	"github.com/signalfx/golib/sfxclient"
	"github.com/signalfx/signalfx-agent/internal/core/common/kubelet"
	"github.com/signalfx/signalfx-agent/internal/core/common/kubernetes"
	"github.com/signalfx/signalfx-agent/internal/core/config"
	"github.com/signalfx/signalfx-agent/internal/monitors"
	"github.com/signalfx/signalfx-agent/internal/monitors/types"
	"github.com/signalfx/signalfx-agent/internal/utils"
	log "github.com/sirupsen/logrus"
	k8s "k8s.io/client-go/kubernetes"
	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

const monitorType = "kubernetes-volumes"

// MONITOR(kubernetes-volumes): This monitor sends usage stats about volumes
// mounted to Kubernetes pods (e.g. free space/inodes).  This information is
// gotten from the Kubelet /stats/summary endpoint.  The normal `collectd/df`
// monitor generally will not report Persistent Volume usage metrics because
// those volumes are not seen by the agent since they can be mounted
// dynamically and older versions of K8s don't support mount propagation of
// those mounts to the agent container.
//
// Dimensions that identify the underlying volume source will be added for
// `awsElasticBlockStore` and `glusterfs` volumes.  Support for more can be
// easily added as needed.

// DIMENSION(volume): The volume name as given in the pod spec under `volumes`
// DIMENSION(kubernetes_pod_uid): The UID of the pod that has this volume
// DIMENSION(kubernetes_pod_name): The name of the pod that has this volume
// DIMENSION(kubernetes_namespace): The namespace of the pod that has this volume

var logger = log.WithFields(log.Fields{"monitorType": monitorType})

func init() {
	monitors.Register(monitorType, func() interface{} { return &Monitor{} }, &Config{})
}

// Config for this monitor
type Config struct {
	config.MonitorConfig
	// Kubelet kubeletClient configuration
	KubeletAPI kubelet.APIConfig `yaml:"kubeletAPI" default:""`
	// Configuration of the Kubernetes API client
	KubernetesAPI kubernetes.APIConfig `yaml:"kubernetesAPI"`
}

// Monitor for K8s volume metrics as reported by kubelet
type Monitor struct {
	Output        types.Output
	cancel        func()
	kubeletClient *kubelet.Client
	k8sClient     *k8s.Clientset
	dimCache      map[string]map[string]string
}

// Configure the monitor and kick off volume metric syncing
func (m *Monitor) Configure(conf *Config) error {
	var err error
	m.kubeletClient, err = kubelet.NewClient(&conf.KubeletAPI, &conf.KubernetesAPI)
	if err != nil {
		return err
	}

	m.k8sClient, err = kubernetes.MakeClient(&conf.KubernetesAPI)
	if err != nil {
		return err
	}

	m.dimCache = make(map[string]map[string]string)

	var ctx context.Context
	ctx, m.cancel = context.WithCancel(context.Background())
	utils.RunOnInterval(ctx, func() {
		dps, err := m.getVolumeMetrics()
		if err != nil {
			logger.WithError(err).Error("Could not get volume metrics")
			return
		}

		for i := range dps {
			m.Output.SendDatapoint(dps[i])
		}
	}, time.Duration(conf.IntervalSeconds)*time.Second)

	return nil
}

func (m *Monitor) getVolumeMetrics() ([]*datapoint.Datapoint, error) {
	summary, err := m.getSummary()
	if err != nil {
		return nil, err
	}

	var dps []*datapoint.Datapoint
	for _, p := range summary.Pods {
		for _, v := range p.VolumeStats {
			dims := map[string]string{
				"volume":               v.Name,
				"kubernetes_pod_uid":   p.PodRef.UID,
				"kubernetes_pod_name":  p.PodRef.Name,
				"kubernetes_namespace": p.PodRef.Namespace,
			}

			volumeDims, err := m.volumeIDDimsForPod(p.PodRef.Name, p.PodRef.Namespace, p.PodRef.UID, v.Name)
			if err != nil {
				logger.WithFields(log.Fields{
					"error":     err,
					"volName":   v.Name,
					"podName":   p.PodRef.Name,
					"namespace": p.PodRef.Namespace,
				}).Error("Could not get volume-specific dimensions")
			} else {
				for k, v := range volumeDims {
					dims[k] = v
				}
			}

			if v.AvailableBytes != nil {
				// uint64 -> int64 conversion should be safe since disk sizes
				// aren't going to get that big for a long time.
				dps = append(dps, sfxclient.Gauge("kubernetes.volume_available_bytes", dims, int64(*v.AvailableBytes)))
			}
			if v.CapacityBytes != nil {
				dps = append(dps, sfxclient.Gauge("kubernetes.volume_capacity_bytes", dims, int64(*v.CapacityBytes)))
			}
		}
	}
	return dps, nil
}

func (m *Monitor) getSummary() (*stats.Summary, error) {
	req, err := m.kubeletClient.NewRequest("POST", "/stats/summary/", nil)
	if err != nil {
		return nil, err
	}

	var summary stats.Summary
	err = m.kubeletClient.DoRequestAndSetValue(req, &summary)
	if err != nil {
		return nil, fmt.Errorf("failed to get summary stats from Kubelet URL %q: %v", req.URL.String(), err)
	}

	return &summary, nil
}

// Shutdown stops the metric sync
func (m *Monitor) Shutdown() {
	if m.cancel != nil {
		m.cancel()
	}
}
