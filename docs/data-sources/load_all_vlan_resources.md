---
page_title: "gigavuecore_load_all_vlan_resources Data Source - gigavuecore"
subcategory: ""
description: |-
  Load vlan resource details
---

# gigavuecore_load_all_vlan_resources Data Source

Load vlan resource details

## Example Usage

```terraform
data "gigavuecore_load_all_vlan_resources" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `adv_vlan_manip_count` (Number, computed) - Total number of Vlan Ids used by Advanced VLAN Manipulation on this box
* `available_vlan_ids` (Attributes List, computed) - VlanIds that are available to be used on this box (see [below for nested schema](#nestedatt--available_vlan_ids))
* `available_vlan_ids_comma_separated` (String, computed) - comma separated VLAN string
* `circuit_tunnel_count` (Number, computed) - Total number of Vlan Ids used by Circuit Tunnel on this box
* `flex_inline_count` (Number, computed) - Total number of Vlan Ids used by Flexible Inline Maps on this box
* `gs_mobility_count` (Number, computed) - Total number of Vlan Ids used by Mobility on this box
* `ip_interface_count` (Number, computed) - Total number of Vlan Ids used by IP Interface on this box
* `other_feature_count` (Number, computed) - Total number of Vlan Ids used by Other Features on this box
* `p_vlan_count` (Number, computed) - Total number of Port Vlan Ids used on this box
* `secure_tunnel_count` (Number, computed) - Total number of Vlan Ids used by Secure Tunnel on this box
* `total_count` (Number, computed) - Total number of Vlan Ids on this box
* `unused_count` (Number, computed) - Number of Vlan Ids available as unused on this box
* `used_count` (Number, computed) - Number of Vlan Ids being used on this box
* `used_vlan_ids` (Attributes List, computed) - VlanIds that are being used on this box (see [below for nested schema](#nestedatt--used_vlan_ids))
* `used_vlan_ids_comma_separated` (String, computed) - comma separated VLAN string

<a id="nestedatt--available_vlan_ids"></a>
### Nested Schema for `available_vlan_ids`

Read-Only:

* `value` (Number) - Represents Vlan ID value
* `value_max` (Number) - Represents maximum Vlan ID in a range and If present, value must be greater than 'value'
<a id="nestedatt--used_vlan_ids"></a>
### Nested Schema for `used_vlan_ids`

Read-Only:

* `value` (Number) - Represents Vlan ID value
* `value_max` (Number) - Represents maximum Vlan ID in a range and If present, value must be greater than 'value'

