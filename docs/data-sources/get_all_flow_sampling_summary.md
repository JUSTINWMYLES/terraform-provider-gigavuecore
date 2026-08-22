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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({gsgroup, num_devices, num_devices_in_sample})), computed)

