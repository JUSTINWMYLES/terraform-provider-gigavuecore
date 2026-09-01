---
page_title: "gigavuecore_deploy_or_save_traffic_policy_graph Action - gigavuecore"
subcategory: ""
description: |-
  Perform the batch operation defined in body. Then deploy the traffic policy graph if it has been deployed or save it
---

# gigavuecore_deploy_or_save_traffic_policy_graph Action

Perform the batch operation defined in body. Then deploy the traffic policy graph if it has been deployed or save it

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_deploy_or_save_traffic_policy_graph" "example" {
  config {
    alias    = "example"
    requests = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic policy graph alias
* `requests` (List of Dynamic, optional) - A list of request


