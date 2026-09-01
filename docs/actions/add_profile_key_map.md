---
page_title: "gigavuecore_add_profile_key_map Action - gigavuecore"
subcategory: ""
description: |-
  Add a new key map entry to the profile
---

# gigavuecore_add_profile_key_map Action

Add a new key map entry to the profile

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_add_profile_key_map" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    hostname   = "example"
    key        = "example"
    rule_id    = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `hostname` (String, required) - hostname or IP address
* `key` (String, required) - SSL key alias
* `rule_id` (Number, optional)


