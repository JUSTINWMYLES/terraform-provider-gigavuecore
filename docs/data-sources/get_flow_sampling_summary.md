---
page_title: "gigavuecore_get_flow_sampling_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Flow Sampling Report Summary
---

# gigavuecore_get_flow_sampling_summary Data Source

Load Flow Sampling Report Summary

## Example Usage

```terraform
data "gigavuecore_get_flow_sampling_summary" "example" {
  alias       = null
  cluster_id  = null
  device_ip   = null
  device_mask = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `device_ip` (String, optional) - deviceIp pattern based active flows
* `device_mask` (String, optional) - deviceIpMask

### Attributes

In addition to all arguments above, the following attributes are exported:

* `gsgroup` (String, computed) - alias of gsgroup
* `num_devices` (Number, computed)
* `num_devices_in_sample` (Number, computed)


