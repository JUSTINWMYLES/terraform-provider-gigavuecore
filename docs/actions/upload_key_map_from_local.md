---
page_title: "gigavuecore_upload_key_map_from_local Action - gigavuecore"
subcategory: ""
description: |-
  Upload Key Map file from local
---

# gigavuecore_upload_key_map_from_local Action

Upload Key Map file from local

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_upload_key_map_from_local" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    file       = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
* `file` (String, required) - key map file to upload


