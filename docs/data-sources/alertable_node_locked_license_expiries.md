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

* `use_db` (Boolean, optional) - if true, use FM database as information provider, else retrieve from node

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `count_` (Number)
* `expiry` (List of String)

