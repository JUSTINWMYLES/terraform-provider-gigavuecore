---
page_title: "gigavuecore_get_all_flow_sampling_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Flow Sampling Report Summary
---

# gigavuecore_get_all_flow_sampling_summary Data Source

Load all Flow Sampling Report Summary

## Example Usage

```terraform
data "gigavuecore_get_all_flow_sampling_summary" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `gsgroup` (String) - alias of gsgroup
* `num_devices` (Number)
* `num_devices_in_sample` (Number)

