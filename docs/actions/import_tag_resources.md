---
page_title: "gigavuecore_import_tag_resources Action - gigavuecore"
subcategory: ""
description: |-
  Upload an csv file from local to import tag resources
---

# gigavuecore_import_tag_resources Action

Upload an csv file from local to import tag resources

## Example Usage

```terraform
action "gigavuecore_import_tag_resources" "example" {
  config {
    input = "example"
    operation = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `input` (String, required) - User uploaded file, only csv format is supported
* `operation` (String, required) - specifies the operation type
