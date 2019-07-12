#monitor template

###signalfx-agent repo/docs



_INSTALL TAB_ (on top)

\<monitor logo>
# \<Monitor name> Installation

Metadata associated with the \<name> collectd plugin can be found _here_. The relevant code for the plugin can be found _here_.

- [Requirements and dependencies](#requirements-and-dependencies)
- [Installation](#installation)
- [Configuration](#configuration)

## Requirements and Dependencies

\<text or table format>

## Installation

Steps to install this monitor are described below. 

Step 1. 

Step 2. 

## Configuration 

| Config option | Required | Type                 | Description                  |
| ------------- | -------- | -------------------- | ---------------------------- |
| `kubeletAPI`  | no       | `object (see below)` | Kubelet client configuration |

The **nested** `kubeletAPI` config object has the following fields:

| Config option | Required | Type     | Description                                                  |
| ------------- | -------- | -------- | ------------------------------------------------------------ |
| `url`         | no       | `string` | URL of the Kubelet instance.  This will default to `https://<current node hostname>:10250` if not provided. |
| `authType`    | no       | `string` | Can be `none` for no auth, `tls` for TLS client cert auth, or `serviceAccount` to use the pod default service account token to authenticate. (**default:** `none`) |
| `skipVerify`  | no       | `bool`   | Whether to skip verification of the Kubelet TLS cert (**default:** `true`) |

### YAML config

Sample YAML config with custom query:

```sh
monitors:
- type: collectd/couchbase
  host: 127.0.0.1
  port: 8091
  collectTarget: "NODE"
  clusterName: "my-cluster"
  username: "user"
  password: "password" 
```

### More Configuration options

In addition to the common configuration options shown [here](https://docs.signalfx.com/en/latest/integrations/agent/monitor-config.html), the following configuration options can be set:

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| item 1 | no | `string` | description
| `item2` | no | `string` | description

## 

_FEATURES TAB_

# Features

- [Description](description)
- [Dashboards](dashboards) (built-in)
- [Infrastructure Navigator views](infrastructure-navigator-views) (built-in)

## Description

Monitors \<what> using the information provided by \<what> which collects metrics from \<what> instances by hitting these endpoints: 
* <link to item>
* <link to item>


For more information on the source data, see <https://...>

Monitor type: \<type name>

Monitor Source Code \<link>

Accepts Endpoints: <yes/no>

Multiple Instances Allowed: <yes/no>

## Dashboards

Built-in dashboard samples in which the default metrics from this monitor display are shown below.

<expandable dashboard screen shot thumbnails>

## Infrastructure Navigator views

Sample built-in Infrastructure Navigator views for this monitor are shown below.

<expandable nav view screen shot thumbnails>



_METRICS TAB_

# Metrics

- [Metrics](#metrics)
- [Dimensions](dimensions)

<nav view screen shots>


## Metrics

In addition to the common default metrics that are described [here](https://docs.signalfx.com/en/latest/integrations/agent/monitor-config.html), the following table shows additional optional metrics available for this monitor.

- Metrics that are marked as Included in the table below are sent by default by the Smart Agent as part of a host-based subscription, and you are not charged for them.

- Metrics that are not marked as Included are custom metrics, such as system or service metrics that you configure the Smart Agent to send outside of the default set of metrics. Your SignalFx subscription allows you to send a certain number of custom metrics.

You may need to add a flag to these metrics. Check the config file for flag requirements.


| Name | Type | Included | Description |
| ---  | ---  | ---    | ---         |
| `name1` | counter | ✔ | Total connections count per broker |
| `name2` | gauge | ✔ | Total number of consumers subscribed to destinations on the broker |
| `name3` | gauge |  | Total number of messages that have been acknowledged from the broker. |


### Custom metric configuration

To collect custom metrics, you must configure your monitor to listen for those metrics and then send those metrics to the agent.

To specify custom metrics, add a _metricsToInclude_ filter to the agent configuration file, as shown in the code snippet below. The snippet lists all available custom metrics. Copy and paste the snippet into your monitor configuration file, then delete any custom metrics that you do not want.

Note that some of the custom metrics require you to set a flag in addition to adding them to the metricsToInclude list. Check the monitor configuration file to see if a flag is required for gathering custom metrics.

```
sh
metricsToInclude:
  - metricNames:
    - name1
    - name2
    monitorType: <name>
```

## Dimensions

The following dimensions may occur on metrics emitted by this monitor. Some dimensions may be specific to certain metrics; other dimensions can be configured. You can add extra dimensions to most metrics. The Common Configuration options page [here](https://docs.signalfx.com/en/latest/integrations/agent/monitor-config.html) also describes how to configure for these extra dimensions. 

| Name | Description |
| ---  | ---         |
| `container_id` | The ID of the running container |
| `container_image` | The container image name |
| `container_name` | The container's name as it appears in the pod spec, the same as container_spec_name but retained for backwards compatibility. |
| `container_spec_name` | The container's name as it appears in the pod spec |





_TROUBLESHOOTING TAB_

# Troubleshooting

- [Verify installation](verify-installation)
- [Troubleshooting monitor operation](troubleshooting-monitor-operations)

  

## Verify Installation

To confirm your installation is installed properly...


## Troubleshooting Monitor Operation

<This is troubleshooting the monitor operation.>






_USAGE TAB_


# Usage 

- [Usage strategy](#usage-strategy)
- [How to](#how-to)
- [Sample code](#sample-code])

## Usage strategy

<discusssion of how to use this monitor for meaningful data display>

## How To

<Examples of how to use the monitor, dashboards, and metrics for a meaningful task.>


## Sample code 

<This is where you can put the sample coding that matches the "how to" section above.>












