---
page_title: "gigavuecore_upload_ssl_trust_store_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload inline SSL trust-store from local file
---

# gigavuecore_upload_ssl_trust_store_file Action

Upload inline SSL trust-store from local file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_upload_ssl_trust_store_file" "example" {
  config {
    cluster_id = "example"
    file       = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `file` (String, required) - file to upload to device


