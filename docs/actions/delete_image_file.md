---
page_title: "gigavuecore_delete_image_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete Image by File Name
---

# gigavuecore_delete_image_file Action

Delete Image by File Name

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_image_file" "example" {
  config {
    file_name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `file_name` (String, required) - Filename of the target image file


