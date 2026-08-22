---
page_title: "gigavuecore_delete_image_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete Image by File Name
---

# gigavuecore_delete_image_file Action

Delete Image by File Name

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
