---
page_title: "gigavuecore_create_ha_group Action - gigavuecore"
subcategory: ""
description: |-
  create HA group
---

# gigavuecore_create_ha_group Action

create HA group

## Example Usage

```terraform
action "gigavuecore_create_ha_group" "example" {
  config {
    fm_ha_tunnel = "example"
    name         = "example"
    nodes        = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `fm_ha_tunnel` (Dynamic, required)
* `name` (String, required) - Name of the new FMHA cluster
* `nodes` (Dynamic, required)


