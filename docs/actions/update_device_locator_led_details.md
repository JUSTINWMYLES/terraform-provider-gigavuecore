---
page_title: "gigavuecore_update_device_locator_led_details Action - gigavuecore"
subcategory: ""
description: |-
  Enable or disable Device Locator Led Details
---

# gigavuecore_update_device_locator_led_details Action

Enable or disable Device Locator Led Details

## Example Usage

```terraform
action "gigavuecore_update_device_locator_led_details" "example" {
  config {
    box_id = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - Box Id of the device
* `cluster_id` (String, required) - Target Cluster ID
