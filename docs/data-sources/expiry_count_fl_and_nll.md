---
page_title: "gigavuecore_expiry_count_fl_and_nll Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns expiry count for floating and node-locked licenses, in that order
---

# gigavuecore_expiry_count_fl_and_nll Data Source

Returns expiry count for floating and node-locked licenses, in that order

## Example Usage

```terraform
data "gigavuecore_expiry_count_fl_and_nll" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Number), computed) - Two elements in the array, first one for floating license expiry count and second one for node-locked licenses expiry count

