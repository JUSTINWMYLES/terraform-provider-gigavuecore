---
page_title: "gigavuecore_redefine_gs_group_sip_ports Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Sip Ports
---

# gigavuecore_redefine_gs_group_sip_ports Action

Redefine GS Group's Sip Ports

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sip_ports" "example" {
  config {
    alias = "example"
    ports = [ 0 ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `ports` (List of Number, optional) - list of TCP/UDP ports. Valid ports 1 - 65535. maximum 10 ports


