---
page_title: "gigavuecore_get_fm_hastatus Data Source - gigavuecore"
subcategory: ""
description: |-
  get status of FmHa
---

# gigavuecore_get_fm_hastatus Data Source

get status of FmHa

## Example Usage

```terraform
data "gigavuecore_get_fm_hastatus" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `role` (String, computed) - for single nodes the role is standby and for clustered nodes the role is active or standalone

