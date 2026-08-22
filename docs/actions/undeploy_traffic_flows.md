---
page_title: "gigavuecore_undeploy_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Undeploy a traffic flow by alias
---

# gigavuecore_undeploy_traffic_flows Action

Undeploy a traffic flow by alias

## Example Usage

```terraform
action "gigavuecore_undeploy_traffic_flows" "example" {
  config {
    alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias
