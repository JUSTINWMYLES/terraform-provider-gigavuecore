---
page_title: "gigavuecore_alertable_node_locked_license_expiries Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns list of expiry info per category for pertinent node-locked licenses, count and a list of details of the expiry, when expiring (expired), card serial number, features, etc
---

# gigavuecore_alertable_node_locked_license_expiries Data Source

Returns list of expiry info per category for pertinent node-locked licenses, count and a list of details of the expiry, when expiring (expired), card serial number, features, etc

## Example Usage

```terraform
data "gigavuecore_alertable_node_locked_license_expiries" "example" {
  use_db = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `use_db` (Bool, optional) - if true, use FM database as information provider, else retrieve from node

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({count_, expiry})), computed)

