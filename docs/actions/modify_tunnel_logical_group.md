---
page_title: "gigavuecore_modify_tunnel_logical_group Action - gigavuecore"
subcategory: ""
description: |-
  Edit Tunnel Logical Group. API handles activate/deactivate tunnel logical group and modification of Tunnel logical group name.
---

# gigavuecore_modify_tunnel_logical_group Action

Edit Tunnel Logical Group. API handles activate/deactivate tunnel logical group and modification of Tunnel logical group name.

## Example Usage

```terraform
action "gigavuecore_modify_tunnel_logical_group" "example" {
  config {
    tunnel_logical_groups = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `tunnel_logical_groups` (List of Dynamic, required)


