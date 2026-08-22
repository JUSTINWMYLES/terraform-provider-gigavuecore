---
page_title: "gigavuecore_get_system_ndp Data Source - gigavuecore"
subcategory: ""
description: |-
  Get System Ndp Refresh Interval
---

# gigavuecore_get_system_ndp Data Source

Get System Ndp Refresh Interval

## Example Usage

```terraform
data "gigavuecore_get_system_ndp" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `ndp_refresh_interval` (Number, computed) - System Ndp Refresh Interval in Seconds

