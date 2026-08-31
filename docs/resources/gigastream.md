---
page_title: "gigavuecore_gigastream Resource - gigavuecore"
subcategory: ""
description: |-
  Find gigastream by alias
---

# gigavuecore_gigastream Resource

Find gigastream by alias

## Example Usage

```terraform
resource "gigavuecore_gigastream" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  drop_weight     = 0
  failover_status = "enable"
  hash_size       = 1
  hash_tool_port = [{
    hash_bucket_ids = [ 0 ]
    tool_ports      = [ "example" ]
  }]
  hash_type    = "advanced"
  hash_weights = [ 0 ]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  ports              = [ "example" ]
  threshold_level    = "Global-Level"
  variance_threshold = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - gigastream alias. Uniquely identifies a gigastream within a cluster
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `drop_weight` (Number, optional) - relative weight for dropping the traffic
* `failover_status` (String, optional) - Failover Status
* `hash_size` (Number, optional) - Hash bucket size
* `hash_tool_port` (Attributes List, optional) - Hash bucket id to tool port mapping (see [below for nested schema](#nestedatt--hash_tool_port))
* `hash_type` (String, optional)
* `hash_weights` (List of Number, optional) - hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `ports` (List of String, optional) - list of the ports to combine into a gigastream
* `threshold_level` (String, optional) - Threshold level
* `variance_threshold` (String, optional) - Variance threshold percentage

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--hash_tool_port"></a>
### Nested Schema for `hash_tool_port`

Required:

* `hash_bucket_ids` (List of Number) - hash bucket id or range
* `tool_ports` (List of String) - tool port(s) mapped to hashBucketIds

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
terraform import gigavuecore_gigastream.example {alias}/{cluster_id}
```
