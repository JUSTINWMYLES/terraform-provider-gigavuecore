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
  alias = null
  cluster_id = null
  comment = null
  drop_weight = null
  failover_status = null
  hash_size = null
  hash_tool_port = []
  hash_type = null
  hash_weights = []
  health_state = null
  health_state_reasons = []
  ports = []
  threshold_level = null
  variance_threshold = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - gigastream alias. Uniquely identifies a gigastream within a cluster
* `cluster_id` (String, optional) - id of the defining cluster
* `comment` (String, optional)
* `drop_weight` (Number, optional) - relative weight for dropping the traffic
* `failover_status` (String, optional) - Failover Status
* `hash_size` (Number, optional) - Hash bucket size
* `hash_tool_port` (List(Object({hash_bucket_ids, tool_ports})), optional) - Hash bucket id to tool port mapping
* `hash_type` (String, optional)
* `hash_weights` (List(Number), optional) - hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `ports` (List(String), optional) - list of the ports to combine into a gigastream
* `threshold_level` (String, optional) - Threshold level
* `variance_threshold` (String, optional) - Variance threshold percentage

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `drop_weight` (Number, computed) - relative weight for dropping the traffic
* `failover_status` (String, computed) - Failover Status
* `hash_size` (Number, computed) - Hash bucket size
* `hash_tool_port` (List(Object({hash_bucket_ids, tool_ports})), computed) - Hash bucket id to tool port mapping
* `hash_type` (String, computed)
* `hash_weights` (List(Number), computed) - hashWeights for 'ports'.If included, the list size must match the size of the 'ports' list
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `ports` (List(String), computed) - list of the ports to combine into a gigastream
* `threshold_level` (String, computed) - Threshold level
* `variance_threshold` (String, computed) - Variance threshold percentage

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_gigastream.example {alias}
```
