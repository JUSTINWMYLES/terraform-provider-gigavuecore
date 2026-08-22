---
page_title: "gigavuecore_upload_key_handler_from_local Action - gigavuecore"
subcategory: ""
description: |-
  Upload world and module file from local
---

# gigavuecore_upload_key_handler_from_local Action

Upload world and module file from local

## Example Usage

```terraform
action "gigavuecore_upload_key_handler_from_local" "example" {
  config {
    alias = "example"
    file = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `file` (String, required) - world or module file to upload
