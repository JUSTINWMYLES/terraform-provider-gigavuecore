---
page_title: "gigavuecore_get_ems_registration_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get status of FM registration with EMS
---

# gigavuecore_get_ems_registration_status Data Source

Get status of FM registration with EMS

## Example Usage

```terraform
data "gigavuecore_get_ems_registration_status" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `error` (String, computed)
* `is_registered` (Boolean, computed)


