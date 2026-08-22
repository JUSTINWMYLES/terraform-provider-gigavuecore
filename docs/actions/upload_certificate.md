---
page_title: "gigavuecore_upload_certificate Action - gigavuecore"
subcategory: ""
description: |-
  Upload a certificate file to push SSL Certificate Global configuration to all the devices
---

# gigavuecore_upload_certificate Action

Upload a certificate file to push SSL Certificate Global configuration to all the devices

## Example Usage

```terraform
action "gigavuecore_upload_certificate" "example" {
  config {
    config = null
    config_level = "example"
    config_level_value = [ "example" ]
    config_type = "example"
    modifiable = true
    ref_count = 1
    template_name = "example"
    update_time = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `config` (Dynamic, optional)
* `config_level` (String, optional) - Scope of the applied FM template
* `config_level_value` (List(String), optional)
* `config_type` (String, optional) - Configuration Type of the FM template
* `modifiable` (Bool, optional)
* `ref_count` (Number, optional)
* `template_name` (String, optional)
* `update_time` (String, optional)
