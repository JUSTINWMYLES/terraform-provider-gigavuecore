---
page_title: "gigavuecore_delete_all_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Delete all traffic flows
---

# gigavuecore_delete_all_traffic_flows Action

Delete all traffic flows

## Example Usage

```terraform
action "gigavuecore_delete_all_traffic_flows" "example" {
  config {
    async = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `async` (Bool, optional) - If true, deletion is processed asynchronously
