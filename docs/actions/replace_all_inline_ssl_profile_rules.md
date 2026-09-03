---
page_title: "gigavuecore_replace_all_inline_ssl_profile_rules Action - gigavuecore"
subcategory: ""
description: |-
  Replace all inline SSL profile rules
---

# gigavuecore_replace_all_inline_ssl_profile_rules Action

Replace all inline SSL profile rules

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_replace_all_inline_ssl_profile_rules" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rules      = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `rules` (List of Dynamic, optional) - inline SSL profile rules


