---
page_title: "gigavuecore_redefine_gs_group_xpkt_match Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Xpkt match
---

# gigavuecore_redefine_gs_group_xpkt_match Action

Redefine GS Group's Xpkt match

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_xpkt_match" "example" {
  config {
    alias   = "example"
    enabled = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `enabled` (Boolean, optional)


