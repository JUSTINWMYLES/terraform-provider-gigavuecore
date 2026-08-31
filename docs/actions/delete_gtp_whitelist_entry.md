---
page_title: "gigavuecore_delete_gtp_whitelist_entry Action - gigavuecore"
subcategory: ""
description: |-
  Delete GTP Whitelist Entry
---

# gigavuecore_delete_gtp_whitelist_entry Action

Delete GTP Whitelist Entry

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_gtp_whitelist_entry" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    imsi       = "example"
    ran        = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID
* `imsi` (String, required) - imsi-based whitelist entry being deleted.
* `ran` (String, optional) - ran-based whitelist entry being queried


