<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# net-io

This monitor reports I/O metrics about network interfaces.

On Linux hosts, this monitor relies on the `/proc` filesystem.
If the underlying host's `/proc` file system is mounted somewhere other than
/proc please specify the path using the top level configuration `procPath`.

```yaml
procPath: /proc
monitors:
 - type: net-io
```


Monitor Type: `net-io`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/netio)

**Accepts Endpoints**: No

**Multiple Instances Allowed**: Yes

## Configuration

**For a list of monitor options that are common to all monitors, see [Common
Configuration](../monitor-config.md#common-configuration).**


| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `interfaces` | no | `list of strings` | The network interfaces to send metrics about. This is an [overridable set](https://docs.signalfx.com/en/latest/integrations/agent/filtering.html#overridable-filters). (**default:** `[* !/^lo\d*$/ !/^docker.*/ !/^t(un|ap)\d*$/ !/^veth.*$/ !/^Loopback*/]`) |


## Metrics

These are the metrics available for this monitor.
Metrics that are categorized as
[container/host](https://docs.signalfx.com/en/latest/admin-guide/usage.html#about-custom-bundled-and-high-resolution-metrics)
(*default*) are ***in bold and italics*** in the list below.


 - ***`if_errors.rx`*** (*cumulative*)<br>    Count of receive errors on the interface
 - ***`if_errors.tx`*** (*cumulative*)<br>    Count of transmit errors on the interface
 - ***`if_octets.rx`*** (*cumulative*)<br>    Count of bytes (octets) received on the interface
 - ***`if_octets.tx`*** (*cumulative*)<br>    Count of bytes (octets) transmitted by the interface
 - `if_packets.rx` (*cumulative*)<br>    Count of packets received on the interface
 - `if_packets.tx` (*cumulative*)<br>    Count of packets transmitted by the interface
 - ***`network.total`*** (*cumulative*)<br>    Total amount of inbound and outbound network traffic on this host, in bytes.  This metric reports with plugin dimension set to "signalfx-metadata".

### Non-default metrics (version 4.7.0+)

**The following information applies to the agent version 4.7.0+ that has
`enableBuiltInFiltering: true` set on the top level of the agent config.**

To emit metrics that are not _default_, you can add those metrics in the
generic monitor-level `extraMetrics` config option.  Metrics that are derived
from specific configuration options that do not appear in the above list of
metrics do not need to be added to `extraMetrics`.

To see a list of metrics that will be emitted you can run `agent-status
monitors` after configuring this monitor in a running agent instance.

### Legacy non-default metrics (version < 4.7.0)

**The following information only applies to agent version older than 4.7.0. If
you have a newer agent and have set `enableBuiltInFiltering: true` at the top
level of your agent config, see the section above. See upgrade instructions in
[Old-style whitelist filtering](../legacy-filtering.md#old-style-whitelist-filtering).**

If you have a reference to the `whitelist.json` in your agent's top-level
`metricsToExclude` config option, and you want to emit metrics that are not in
that whitelist, then you need to add an item to the top-level
`metricsToInclude` config option to override that whitelist (see [Inclusion
filtering](../legacy-filtering.md#inclusion-filtering).  Or you can just
copy the whitelist.json, modify it, and reference that in `metricsToExclude`.



