---
page_title: "gigavuecore_delete_all_gtp_whitelist_entries Action - gigavuecore"
subcategory: ""
description: |-
  Delete All GTP Whitelist Entries
---

# gigavuecore_delete_all_gtp_whitelist_entries Action

Delete All GTP Whitelist Entries

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_gtp_whitelist_entries" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID


