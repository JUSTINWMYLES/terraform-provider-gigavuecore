---
page_title: "gigavuecore_stack_link Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new StackLink
---

# gigavuecore_stack_link Resource

Create a new StackLink

## Example Usage

```terraform
resource "gigavuecore_stack_link" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  endpoint1    = "example"
  endpoint2    = "example"
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  port_list1 = ["example"]
  port_list2 = ["example"]
  type       = "gigastream"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Port Pair alias. Uniquely identifies a StackLink within a cluster
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `endpoint1` (String, required) - for 'gigastream' type, this is the gigastream alias of node1. for 'port' type, this is the port id on node1
* `endpoint2` (String, required) - for 'gigastream' type, this is the gigastream alias of node2. for 'port' type, this is the port id on node2
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `port_list1` (List of String, optional) - only applicable for type 'gigastream'. port-list corresponding to endpoint1
* `port_list2` (List of String, optional) - only applicable for type 'gigastream'. port-list corresponding to endpoint2
* `type` (String, required) - specifies whether this is a gigastream-based or port-based stack link

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_stack_link.example {alias}/{cluster_id}
```
