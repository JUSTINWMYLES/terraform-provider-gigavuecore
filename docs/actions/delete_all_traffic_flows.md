---
page_title: "gigavuecore_delete_all_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Delete all traffic flows
---

# gigavuecore_delete_all_traffic_flows Action

Delete all traffic flows

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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

* `async` (Boolean, optional) - If true, deletion is processed asynchronously


