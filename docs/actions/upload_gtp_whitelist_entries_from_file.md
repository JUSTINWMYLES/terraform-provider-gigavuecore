---
page_title: "gigavuecore_upload_gtp_whitelist_entries_from_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload GTP Whitelist Entries from file
---

# gigavuecore_upload_gtp_whitelist_entries_from_file Action

Upload GTP Whitelist Entries from file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_upload_gtp_whitelist_entries_from_file" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    entries    = "example"
    usage      = "create"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID
* `entries` (String, required) - File containing GTP Whitelist Entries (one entry per line)
* `usage` (String, required) - When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identified by the 'alias' parameter


