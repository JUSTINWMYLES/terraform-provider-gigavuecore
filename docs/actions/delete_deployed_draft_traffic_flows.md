---
page_title: "gigavuecore_delete_deployed_draft_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Delete deployed draft Traffic Flows by alias
---

# gigavuecore_delete_deployed_draft_traffic_flows Action

Delete deployed draft Traffic Flows by alias

## Example Usage

```terraform
action "gigavuecore_delete_deployed_draft_traffic_flows" "example" {
  config {
    alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic Flows alias or ID
