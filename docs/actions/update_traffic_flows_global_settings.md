---
page_title: "gigavuecore_update_traffic_flows_global_settings Action - gigavuecore"
subcategory: ""
description: |-
  Update global settings for Traffic Flows
---

# gigavuecore_update_traffic_flows_global_settings Action

Update global settings for Traffic Flows

## Example Usage

```terraform
action "gigavuecore_update_traffic_flows_global_settings" "example" {
  config {
    body = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body` (Dynamic, required) - Settings for traffic flows, including migration and fabric resource configuration.


