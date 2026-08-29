---
page_title: "gigavuecore_update_device_port Action - gigavuecore"
subcategory: ""
description: |-
  Update Device Port configuration
---

# gigavuecore_update_device_port Action

Update Device Port configuration

## Example Usage

```terraform
action "gigavuecore_update_device_port" "example" {
  config {
    admin_status  = "example"
    alias         = "example"
    auto_neg      = true
    body_port_id  = "example"
    breakout_mode = "example"
    cable_length  = "example"
    cluster_id    = "example"
    comment       = "example"
    config_speed  = "example"
    duplex        = "example"
    force_link_up = true
    mtu           = 0
    port_id       = "example"
    port_type     = "example"
    ude = {
      enabled = true
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `admin_status` (String, optional)
* `alias` (String, optional) - device port alias
* `auto_neg` (Boolean, optional)
* `body_port_id` (String, required) - device port id. used to identify target port.
* `breakout_mode` (String, optional) - 4x = 4x10G; 2q = 2x40G
* `cable_length` (String, optional) - Attached cable length in meter
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `config_speed` (String, optional)
* `duplex` (String, optional)
* `force_link_up` (Boolean, optional)
* `mtu` (Number, optional)
* `port_id` (String, required) - Device Port ID (format: boxId\_slotId\_port, example: 1\_1\_c1)
* `port_type` (String, optional)
* `ude` (Attributes, optional) - Unidirectional Ethernet (see [below for nested schema](#nestedatt--ude))

<a id="nestedatt--ude"></a>
### Nested Schema for `ude`

Optional:

* `enabled` (Boolean) - Only applicable if 100g-bidi is detected

