---
page_title: "gigavuecore_get_all_cluster_licenses Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all clusters with all the chassis and cards information and licenses installed in  Hierarchical View
---

# gigavuecore_get_all_cluster_licenses Data Source

Get all clusters with all the chassis and cards information and licenses installed in  Hierarchical View

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_licenses" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `chassis_modules` (Attributes List) (see [below for nested schema](#nestedatt--items--chassis_modules))
* `cluster_name` (String)
* `hash_code` (String)
<a id="nestedatt--items--chassis_modules"></a>
### Nested Schema for `items.chassis_modules`

Read-Only:

* `box_id` (Number)
* `cards` (Attributes List) (see [below for nested schema](#nestedatt--items--chassis_modules--cards))
* `cluster_name` (String)
* `host_name` (String)
* `hw_type` (String)
* `licenses` (Attributes List) (see [below for nested schema](#nestedatt--items--chassis_modules--licenses))
* `model` (String) - Gigamon physical device models
* `node_id` (String)
* `serial_number` (String)
<a id="nestedatt--items--chassis_modules--cards"></a>
### Nested Schema for `items.chassis_modules.cards`

Read-Only:

* `box_id` (Number)
* `cluster_name` (String) - name of the cluster of chassis
* `hw_type` (String) - the hardware type for the card
* `licenses` (Attributes List) (see [below for nested schema](#nestedatt--items--chassis_modules--cards--licenses))
* `serial_number` (String)
* `slot_id` (String) - chassis slot number where the card is inserted
<a id="nestedatt--items--chassis_modules--cards--licenses"></a>
### Nested Schema for `items.chassis_modules.cards.licenses`

Read-Only:

* `end_date` (Number) - The license end date code in format YYYYMMDD
* `features` (String) - Comma-separated list of licensed application names
* `floated` (Boolean) - true if the license was assigned using current Fabric Manager from a floating license pool, false otherwise
* `grace_period` (Number) - number of days after the end date of the license for which the user can still use the license as if it was still valid
* `license_status` (String) - status of the license
* `license_type` (String) - type of the license
* `notif_period` (Number) - number of days before the end date of the license when notification is sent out about impending license expiry
* `start_date` (Number) - The license start date code in format YYYYMMDD
<a id="nestedatt--items--chassis_modules--licenses"></a>
### Nested Schema for `items.chassis_modules.licenses`

Read-Only:

* `end_date` (Number) - The license end date code in format YYYYMMDD
* `features` (String) - Comma-separated list of licensed application names
* `floated` (Boolean) - true if the license was assigned using current Fabric Manager from a floating license pool, false otherwise
* `grace_period` (Number) - number of days after the end date of the license for which the user can still use the license as if it was still valid
* `license_status` (String) - status of the license
* `license_type` (String) - type of the license
* `notif_period` (Number) - number of days before the end date of the license when notification is sent out about impending license expiry
* `start_date` (Number) - The license start date code in format YYYYMMDD

