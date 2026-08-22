---
page_title: "gigavuecore_inline_network_group Resource - gigavuecore"
subcategory: ""
description: |-
  Find Inline Network Group by alias
---

# gigavuecore_inline_network_group Resource

Find Inline Network Group by alias

## Example Usage

```terraform
resource "gigavuecore_inline_network_group" "example" {
  alias = null
  bundled = null
  comment = null
  health_state = null
  health_state_reasons = []
  inline_networks = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Tool Group alias. Unique within a cluster
* `bundled` (Bool, optional) - Bundled: an inline tool or group of inline tools is injected into a link bundle between two networks; Distinct: an inline tool or group of inline tools is shared by a number of pairs of networks
* `comment` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), optional)
* `inline_networks` (Set(String), required) - list of inline-network aliases

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `bundled` (Bool, computed) - Bundled: an inline tool or group of inline tools is injected into a link bundle between two networks; Distinct: an inline tool or group of inline tools is shared by a number of pairs of networks
* `comment` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `id` (String, computed)

