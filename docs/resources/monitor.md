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
  alias = "example"
  cache = {
    export_triggers = {
      event            = "example"
      timeout_active   = 0
      timeout_inactive = 0
    }
    type = "example"
  }
  cluster_id  = "example"
  description = "example"
  records     = [ "example" ]
  sampling = {
    mode                 = "example"
    single_sampling_rate = 0
  }
  sampling_space = 0
  ssl_port_restrictions = {
    ports     = [ 0 ]
    ssl_ports = "example"
  }
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cache` (Attributes, required) - Netflow Monitor Cache (see [below for nested schema](#nestedatt--cache))
* `cluster_id` (String, required) - id of the defining cluster
* `description` (String, optional)
* `records` (Set of String, optional) - Aliases of referenced Netflow Records defined on the GsGroup. Up to 5 records.
* `sampling` (Attributes, optional) - monitor sampling (see [below for nested schema](#nestedatt--sampling))
* `sampling_space` (Number, optional) - DEPRECATED: use 'sampling'
* `ssl_port_restrictions` (Attributes, optional) - Port restrictions for Netflow/SSL sessions (see [below for nested schema](#nestedatt--ssl_port_restrictions))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--cache"></a>
### Nested Schema for `cache`

Required:

* `export_triggers` (Attributes) - Netflow Monitor Cache Export Triggers (see [below for nested schema](#nestedatt--cache--export_triggers))

Optional:

* `type` (String)

<a id="nestedatt--cache--export_triggers"></a>
### Nested Schema for `cache.export_triggers`

Optional:

* `event` (String)
* `timeout_active` (Number) - in seconds. max value is 7 days. default is 30 min
* `timeout_inactive` (Number) - in seconds. max value is 7 days. default is 15 sec

<a id="nestedatt--sampling"></a>
### Nested Schema for `sampling`

Optional:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)

<a id="nestedatt--ssl_port_restrictions"></a>
### Nested Schema for `ssl_port_restrictions`

Optional:

* `ports` (List of Number) - The list of TCP ports whose packets will be sent to the SSL module. Not valid if sslPort is 'all'. if 'wellKnownPorts' is selected for 'sslPorts' ports will be \[993, 995, 465, 636, 563, 443\]
* `ssl_ports` (String) - Ports whose packets will be sent to the SSL module. 'all' ports; 'wellKnownPorts' is \[993, 995, 465, 636, 563, 443\]; or 'ports' - specify upto 10 ports
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_monitor.example {alias}/{cluster_id}
```
