---
page_title: "gigavuecore_import_tag_resources Action - gigavuecore"
subcategory: ""
description: |-
  Upload an csv file from local to import tag resources
---

# gigavuecore_import_tag_resources Action

Upload an csv file from local to import tag resources

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_import_tag_resources" "example" {
  config {
    input     = "example"
    operation = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `input` (String, required) - User uploaded file, only csv format is supported
* `operation` (String, required) - specifies the operation type


