---
page_title: "gigavuecore_upgrade_fm_image Action - gigavuecore"
subcategory: ""
description: |-
  Upgrade FM image
---

# gigavuecore_upgrade_fm_image Action

Upgrade FM image

## Example Usage

```terraform
action "gigavuecore_upgrade_fm_image" "example" {
  config {
    async        = true
    file_path    = "example"
    image_server = "example"
    reboot       = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `async` (Boolean, optional) - if provided, the call returns immediately with the \[202 Accepted\] HTTP status code, and the image upgrade will complete in the background
* `file_path` (String, required) - image file path on the image server
* `image_server` (String, required) - alias of an image file server. has to reference one of the existing image file server profiles
* `reboot` (Boolean, required) - indicates whether nodes should reboot after image upgrade


