---
page_title: "gigavuecore_delete_deployed_draft_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Delete deployed draft Traffic Flows by alias
---

# gigavuecore_delete_deployed_draft_traffic_flows Action

Delete deployed draft Traffic Flows by alias

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


