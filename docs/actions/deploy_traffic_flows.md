---
page_title: "gigavuecore_deploy_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Deploy a Traffic Flows configuration
---

# gigavuecore_deploy_traffic_flows Action

Deploy a Traffic Flows configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_deploy_traffic_flows" "example" {
  config {
    alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic Flows alias or ID


