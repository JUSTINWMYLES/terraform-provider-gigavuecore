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
  alias                = null
  bundled              = null
  comment              = null
  health_state         = null
  health_state_reasons = []
  inline_networks      = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Tool Group alias. Unique within a cluster
* `bundled` (Boolean, optional) - Bundled: an inline tool or group of inline tools is injected into a link bundle between two networks; Distinct: an inline tool or group of inline tools is shared by a number of pairs of networks
* `comment` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_networks` (Set of String, required) - list of inline-network aliases

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `bundled` (Boolean, computed) - Bundled: an inline tool or group of inline tools is injected into a link bundle between two networks; Distinct: an inline tool or group of inline tools is shared by a number of pairs of networks
* `comment` (String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `id` (String, computed)

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

