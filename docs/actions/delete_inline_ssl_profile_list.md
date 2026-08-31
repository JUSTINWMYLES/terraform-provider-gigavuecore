---
page_title: "gigavuecore_delete_inline_ssl_profile_list Action - gigavuecore"
subcategory: ""
description: |-
  Delete inline SSL profile nodecryptlist or decryptlist
---

# gigavuecore_delete_inline_ssl_profile_list Action

Delete inline SSL profile nodecryptlist or decryptlist

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_inline_ssl_profile_list" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    list_type  = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `list_type` (String, required) - specify a nodecryptlist or decryptlist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)


