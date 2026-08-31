---
page_title: "gigavuecore_redefine_inline_network_group Action - gigavuecore"
subcategory: ""
description: |-
  Redefine an Inline Network Group
---

# gigavuecore_redefine_inline_network_group Action

Redefine an Inline Network Group

## Example Usage

```terraform
action "gigavuecore_redefine_inline_network_group" "example" {
  config {
    alias        = "example"
    body_alias   = "example"
    bundled      = true
    cluster_id   = "example"
    comment      = "example"
    health_state = "green"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "green"
      traffic_health_state_computation_type = "PORT_LOW_UTIL"
    }]
    inline_networks = [ "example" ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of Inline Network Group to redefine
* `body_alias` (String, required) - Inline Tool Group alias. Unique within a cluster
* `bundled` (Boolean, optional) - Bundled: an inline tool or group of inline tools is injected into a link bundle between two networks; Distinct: an inline tool or group of inline tools is shared by a number of pairs of networks
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, optional) (see [below for nested schema](#nestedatt--health_state_reasons))
* `inline_networks` (Set of String, required) - list of inline-network aliases

<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

