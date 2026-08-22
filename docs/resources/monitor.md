---
page_title: "gigavuecore_monitor Resource - gigavuecore"
subcategory: ""
description: |-
  Find Netflow Monitor by alias
---

# gigavuecore_monitor Resource

Find Netflow Monitor by alias

## Example Usage

```terraform
resource "gigavuecore_monitor" "example" {
  alias = null
  cache = {}
  cluster_id = null
  description = null
  records = []
  sampling = {}
  sampling_space = null
  ssl_port_restrictions = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cache` (Object({export_triggers, type}), required) - Netflow Monitor Cache
  * `export_triggers` (Object({event, timeout_active, timeout_inactive}), required) - Netflow Monitor Cache Export Triggers
    * `event` (String, optional)
    * `timeout_active` (Number, optional) - in seconds. max value is 7 days. default is 30 min
    * `timeout_inactive` (Number, optional) - in seconds. max value is 7 days. default is 15 sec
  * `type` (String, optional)
* `cluster_id` (String, optional) - id of the defining cluster
* `description` (String, optional)
* `records` (Set(String), optional) - Aliases of referenced Netflow Records defined on the GsGroup. Up to 5 records.
* `sampling` (Object({mode, single_sampling_rate}), optional) - monitor sampling
  * `mode` (String, optional)
  * `single_sampling_rate` (Number, optional) - Packet interval window size. Valid values: 10-16000 (in packets)
* `sampling_space` (Number, optional) - DEPRECATED: use 'sampling'
* `ssl_port_restrictions` (Object({ports, ssl_ports}), optional) - Port restrictions for Netflow/SSL sessions
  * `ports` (List(Number), optional) - The list of TCP ports whose packets will be sent to the SSL module. Not valid if sslPort is 'all'. if 'wellKnownPorts' is selected for 'sslPorts' ports will be \[993, 995, 465, 636, 563, 443\]
  * `ssl_ports` (String, optional) - Ports whose packets will be sent to the SSL module. 'all' ports; 'wellKnownPorts' is \[993, 995, 465, 636, 563, 443\]; or 'ports' - specify upto 10 ports

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `description` (String, computed)
* `records` (Set(String), computed) - Aliases of referenced Netflow Records defined on the GsGroup. Up to 5 records.
* `sampling` (Object({mode, single_sampling_rate}), computed) - monitor sampling
  * `mode` (String, optional)
  * `single_sampling_rate` (Number, optional) - Packet interval window size. Valid values: 10-16000 (in packets)
* `sampling_space` (Number, computed) - DEPRECATED: use 'sampling'
* `ssl_port_restrictions` (Object({ports, ssl_ports}), computed) - Port restrictions for Netflow/SSL sessions
  * `ports` (List(Number), optional) - The list of TCP ports whose packets will be sent to the SSL module. Not valid if sslPort is 'all'. if 'wellKnownPorts' is selected for 'sslPorts' ports will be \[993, 995, 465, 636, 563, 443\]
  * `ssl_ports` (String, optional) - Ports whose packets will be sent to the SSL module. 'all' ports; 'wellKnownPorts' is \[993, 995, 465, 636, 563, 443\]; or 'ports' - specify upto 10 ports

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_monitor.example {alias}
```
