---
page_title: "gigavuecore_upload_device_config_file Action - gigavuecore"
subcategory: ""
description: |-
  upload device config file
---

# gigavuecore_upload_device_config_file Action

upload device config file

## Example Usage

```terraform
action "gigavuecore_upload_device_config_file" "example" {
  config {
    file = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `file` (String, required) - Attached Config file. In text format
