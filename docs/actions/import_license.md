---
page_title: "gigavuecore_import_license Action - gigavuecore"
subcategory: ""
description: |-
  Import FM License
---

# gigavuecore_import_license Action

Import FM License

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_import_license" "example" {
  config {
    lic_file_name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `lic_file_name` (String, required) - User uploaded license file


