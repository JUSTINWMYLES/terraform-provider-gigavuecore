---
page_title: "gigavuecore_redefine_gs_group_diameter_packet Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Diameter packet
---

# gigavuecore_redefine_gs_group_diameter_packet Action

Redefine GS Group's Diameter packet

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_diameter_packet" "example" {
  config {
    alias   = "example"
    timeout = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `timeout` (Number, optional) - timeout in seconds to remove OOO or fragmented packet


