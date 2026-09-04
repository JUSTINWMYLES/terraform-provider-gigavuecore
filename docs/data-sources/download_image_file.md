---
page_title: "gigavuecore_download_image_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Download Image File
---

# gigavuecore_download_image_file Data Source

Download Image File

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_download_image_file" "example" {
  file_name = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `file_name` (String, required) - Filename of the target image file


