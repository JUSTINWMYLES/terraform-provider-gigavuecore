---
page_title: "gigavuecore_upload_image_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload an image file from local file
---

# gigavuecore_upload_image_file Action

Upload an image file from local file

## Example Usage

```terraform
action "gigavuecore_upload_image_file" "example" {
  config {
    image = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `image` (String, required) - User uploaded file


