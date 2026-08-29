---
page_title: "gigavuecore_stack_link Resource - gigavuecore"
subcategory: ""
description: |-
  Find Stack Link by alias
---

# gigavuecore_stack_link Resource

Find Stack Link by alias

## Example Usage

```terraform
resource "gigavuecore_stack_link" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  endpoint1    = "example"
  endpoint2    = "example"
  health_state = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  port_list1 = [ "example" ]
  port_list2 = [ "example" ]
  type       = "example"
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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_stack_link.example {alias}/{cluster_id}
```
