---
page_title: "gigavuecore_delete_profile_decrypt_port_map_by_id Action - gigavuecore"
subcategory: ""
description: |-
  Delete a port map entry from the profile
---

# gigavuecore_delete_profile_decrypt_port_map_by_id Action

Delete a port map entry from the profile

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_profile_decrypt_port_map_by_id" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rule_id    = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `rule_id` (Number, required) - ruleId of the port map ports to delete


