---
page_title: "gigavuecore_get_all_devices Data Source - gigavuecore"
subcategory: ""
description: |-
  Get a list of physical devices on which license can be allocated, the selection criteria for devices determined by the query parameters
---

# gigavuecore_get_all_devices Data Source

Get a list of physical devices on which license can be allocated, the selection criteria for devices determined by the query parameters

## Example Usage

```terraform
data "gigavuecore_get_all_devices" "example" {
  activation_id = null
  device_model  = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, optional) - activation ID of the license; if provided, devices that are relevant for the SKU / features of the activation are returned
* `device_model` (String, required) - model of device, used only if 'activationId' is null

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (Number)
* `cards` (Attributes List) (see [below for nested schema](#nestedatt--items--cards))
* `cluster_name` (String)
* `deleted` (Boolean)
* `healthy` (Boolean)
* `host_id` (String)
* `host_name` (String)
* `model` (String) - Gigamon physical device models
* `node_id` (String)
* `product_version` (String)
* `serial_number` (String)
<a id="nestedatt--items--cards"></a>
### Nested Schema for `items.cards`

Read-Only:

* `box_id` (Number)
* `cluster_name` (String)
* `healthy` (Boolean)
* `hw_type` (String)
* `serial_number` (String)
* `slot_id` (String)

