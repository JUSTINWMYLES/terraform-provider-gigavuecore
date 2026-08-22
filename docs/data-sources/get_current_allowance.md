---
page_title: "gigavuecore_get_current_allowance Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the daily allowance per current set of Volume-Based Licenses
---

# gigavuecore_get_current_allowance Data Source

Gives the daily allowance per current set of Volume-Based Licenses

## Example Usage

```terraform
data "gigavuecore_get_current_allowance" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `any` (Map(Number), computed)
* `app` (Map(Number), computed)
* `bundle` (Map(Number), computed)

