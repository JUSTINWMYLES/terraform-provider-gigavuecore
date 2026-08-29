---
page_title: "gigavuecore_redefine_gs_group_sip_media Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Media
---

# gigavuecore_redefine_gs_group_sip_media Action

Redefine GS Group's Sip Media

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_media" "example" {
  config {
    alias   = "example"
    timeout = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `timeout` (Number, optional) - Sip media timeout value in seconds .Valid values: 30-300.


