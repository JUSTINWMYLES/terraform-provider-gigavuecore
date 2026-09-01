---
page_title: "gigavuecore_deploy_traffic_policy_graph Action - gigavuecore"
subcategory: ""
description: |-
  Deploy the traffic policy graph
---

# gigavuecore_deploy_traffic_policy_graph Action

Deploy the traffic policy graph

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_deploy_traffic_policy_graph" "example" {
  config {
    alias    = "example"
    requests = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic Policy Graph alias
* `requests` (List of String, optional) - A list of deployment Ids


