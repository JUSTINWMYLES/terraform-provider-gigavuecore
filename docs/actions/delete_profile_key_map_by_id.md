---
page_title: "gigavuecore_delete_profile_key_map_by_id Action - gigavuecore"
subcategory: ""
description: |-
  Delete a key map entry from the profile
---

# gigavuecore_delete_profile_key_map_by_id Action

Delete a key map entry from the profile

## Example Usage

```terraform
action "gigavuecore_delete_profile_key_map_by_id" "example" {
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
* `rule_id` (Number, required) - ruleId of the key map to delete


