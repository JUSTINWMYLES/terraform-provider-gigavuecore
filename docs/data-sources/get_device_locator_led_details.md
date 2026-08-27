---
page_title: "gigavuecore_get_device_locator_led_details Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Device Locator Led Details
---

# gigavuecore_get_device_locator_led_details Data Source

Get Device Locator Led Details

## Example Usage

```terraform
data "gigavuecore_get_device_locator_led_details" "example" {
  box_id     = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - Box Id of the device
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `locator_led` (String, computed)


