---
page_title: "gigavuecore_redefine_gs_group_rtp_ports Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Rtp Ports
---

# gigavuecore_redefine_gs_group_rtp_ports Action

Redefine GS Group's Rtp Ports

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_rtp_ports" "example" {
  config {
    alias = "example"
    range = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `range` (Dynamic, optional) - GsGroup Rtp Port Range Parameters
