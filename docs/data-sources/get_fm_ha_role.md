---
page_title: "gigavuecore_get_fm_ha_role Data Source - gigavuecore"
subcategory: ""
description: |-
  FM HA role
---

# gigavuecore_get_fm_ha_role Data Source

FM HA role

## Example Usage

```terraform
data "gigavuecore_get_fm_ha_role" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `role` (String, computed) - for single nodes the role is standby and for clustered nodes the role is active or standalone

