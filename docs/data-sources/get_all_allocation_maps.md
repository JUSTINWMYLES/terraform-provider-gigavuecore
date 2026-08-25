---
page_title: "gigavuecore_get_all_allocation_maps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get inventory of licenses assigned on the physical devices monitored by the FM
---

# gigavuecore_get_all_allocation_maps Data Source

Get inventory of licenses assigned on the physical devices monitored by the FM

## Example Usage

```terraform
data "gigavuecore_get_all_allocation_maps" "example" {
  encode            = null
  exclude_host_name = null
  exclude_ip        = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `encode` (Boolean, optional) - encode the JSON output
* `exclude_host_name` (String, optional) - exclude the host name in the inventory output
* `exclude_ip` (String, optional) - exclude the host IP in the inventory output

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `chassis` (Attributes) (see [below for nested schema](#nestedatt--items--chassis))
* `deployment_mac` (String)
* `slot_to_licenses` (Attributes List) (see [below for nested schema](#nestedatt--items--slot_to_licenses))
<a id="nestedatt--items--chassis"></a>
### Nested Schema for `items.chassis`

Read-Only:

* `box_id` (Number)
* `cards` (Attributes List) (see [below for nested schema](#nestedatt--items--chassis--cards))
* `cluster_name` (String)
* `deleted` (Boolean)
* `healthy` (Boolean)
* `host_id` (String)
* `host_name` (String)
* `model` (String) - Gigamon physical device models
* `node_id` (String)
* `product_version` (String)
* `serial_number` (String)
<a id="nestedatt--items--chassis--cards"></a>
### Nested Schema for `items.chassis.cards`

Read-Only:

* `box_id` (Number)
* `cluster_name` (String)
* `healthy` (Boolean)
* `hw_type` (String)
* `serial_number` (String)
* `slot_id` (String)
<a id="nestedatt--items--slot_to_licenses"></a>
### Nested Schema for `items.slot_to_licenses`

Read-Only:

* `skus` (Attributes List) (see [below for nested schema](#nestedatt--items--slot_to_licenses--skus))
* `slot_id` (Number)
<a id="nestedatt--items--slot_to_licenses--skus"></a>
### Nested Schema for `items.slot_to_licenses.skus`

Read-Only:

* `end_date` (String)
* `license_key` (String)
* `sku` (String)

