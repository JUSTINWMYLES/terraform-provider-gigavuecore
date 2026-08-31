---
page_title: "gigavuecore_import_tags Action - gigavuecore"
subcategory: ""
description: |-
  Upload an csv file from local  to import tags
---

# gigavuecore_import_tags Action

Upload an csv file from local  to import tags

## Example Usage

```terraform
action "gigavuecore_import_tags" "example" {
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


