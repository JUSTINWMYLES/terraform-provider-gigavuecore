---
page_title: "gigavuecore_configure_device_chassis Action - gigavuecore"
subcategory: ""
description: |-
  Configure Device Chassis
---

# gigavuecore_configure_device_chassis Action

Configure Device Chassis

## Example Usage

```terraform
action "gigavuecore_configure_device_chassis" "example" {
  config {
    box_id       = 1
    chassis_type = "example"
    cluster_id   = "example"
    gdp          = true
    l2_gre_id    = 1
    leaf_config = {
      mode = "default"
    }
    mode          = "100G"
    node_id       = "example"
    serial_number = "example"
    vxlan_id      = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, required) - Configure the chassis Box ID value
* `chassis_type` (String, optional) - Specify chassis type (to provision offline) valid values: hb1, hc1, hc2, hc2-v2, hc3, hd4-ccv1, hd4-ccv2, hd8-ccv1, hd8-ccv2, ly2r, ta1, ta10, ta10a, ta40, itac, tacx
* `cluster_id` (String, required) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `gdp` (Boolean, optional) - enable/disable GDP for chassis
* `l2_gre_id` (Number, optional) - for type l2gre maximum value is 4294967295 , Value of 0 disables the vxlanId
* `leaf_config` (Attributes, optional) - Configuration of Leaf node for Spine-Link (see [below for nested schema](#nestedatt--leaf_config))
* `mode` (String, optional) - Use of 100G ports requires chassis mode value to be 100G. 100G applicable to HC2-v2 only. 100GLeft, Right refers to left, right side chassis bank.
* `node_id` (String, optional) - ID of the target device
* `serial_number` (String, optional) - Specify Chassis Serial Number (Defaults to the local chassis)
* `vxlan_id` (Number, optional) - for type vxlan maximum value is 16777215 , Value of 0 disables the vxlanId

<a id="nestedatt--leaf_config"></a>
### Nested Schema for `leaf_config`

Optional:

* `mode` (String)

