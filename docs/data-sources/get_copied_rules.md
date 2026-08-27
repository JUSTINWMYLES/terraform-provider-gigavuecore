---
page_title: "gigavuecore_get_copied_rules Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve previously copied rules from the user-specific clipboard filtered by rule category and type. Returns the complete set of copied rules matching the specified category and type.
---

# gigavuecore_get_copied_rules Data Source

Retrieve previously copied rules from the user-specific clipboard filtered by rule category and type.
Returns the complete set of copied rules matching the specified category and type.

## Example Usage

```terraform
data "gigavuecore_get_copied_rules" "example" {
  rule_category = null
  rule_type     = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `rule_category` (String, required) - Rule category (SOURCE or APPLICATION)
* `rule_type` (String, required) - Rule type (MapSubType enum value)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `copied_rules` (Attributes, computed) (see [below for nested schema](#nestedatt--copied_rules))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--copied_rules"></a>
### Nested Schema for `copied_rules`

Read-Only:

* `created_time` (Number) - Timestamp when rules were copied
* `policy_id` (String) - MongoDB document ID
* `rule_category` (String) - Category of rules to copy/paste
* `rule_set` (Attributes List) - Set of copied rules (see [below for nested schema](#nestedatt--copied_rules--rule_set))
* `rule_type` (String) - Type of application rule
* `username` (String) - Username of the user who copied the rules
<a id="nestedatt--copied_rules--rule_set"></a>
### Nested Schema for `copied_rules.rule_set`

Read-Only:

* `application_rules` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules))
* `flow_alias` (String) - Flow alias (for APPLICATION category)
* `policy_alias` (String) - Policy alias
* `policy_id` (String) - Policy ID
* `source_and_rule_alias` (String) - Source and rule alias (for SOURCE category)
* `source_rules` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules))
* `sub_flow_alias` (String) - Sub-flow alias (for APPLICATION category)
<a id="nestedatt--copied_rules--rule_set--application_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules`

Read-Only:

* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--gs_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_rules--drop_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_rules.drop_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_rules.drop_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_rules.pass_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_rules.pass_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample5_g_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample5_g_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample5_g_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample5_g_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `priority` (Number)
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample5_g_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_diameter_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_diameter_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_diameter_rules.pass_rules`

Read-Only:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_diameter_rules.pass_rules.diameter`

Read-Only:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_overlap_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_sip_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_sip_rules.pass_rules`

Read-Only:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_sip_rules.pass_rules.sip`

Read-Only:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_sample_sip_rules.pass_rules.sip.id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist5_g_overlap_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist5_g_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist5_g_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_overlap_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_overlap_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--copied_rules--rule_set--application_rules--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.flow_whitelist_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--copied_rules--rule_set--application_rules--gs_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.gs_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--application_rules--gs_rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--application_rules--gs_rules--drop_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.gs_rules.drop_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--application_rules--gs_rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.application_rules.gs_rules.pass_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--copied_rules--rule_set--source_rules"></a>
### Nested Schema for `copied_rules.rule_set.source_rules`

Read-Only:

* `alias` (String) - unique sourcesAndRules alias
* `components` (Attributes List) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--components))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rewrite))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rule_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline' maps; 'passAll' is applicable to 'regular' and 'inline' maps
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules))
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--tags))
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--vlan_tag))
<a id="nestedatt--copied_rules--rule_set--source_rules--components"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.components`

Read-Only:

* `cluster_id` (String) - id of the defining cluster
* `components` (Attributes List) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--components--components))
<a id="nestedatt--copied_rules--rule_set--source_rules--components--components"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.components.components`

Read-Only:

* `ids` (List of Dynamic)
* `type` (String)
<a id="nestedatt--copied_rules--rule_set--source_rules--ip_rewrite"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--copied_rules--rule_set--source_rules--rewrite"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--copied_rules--rule_set--source_rules--rules"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules))
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules--vlan_tag))
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules--rewrite"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules--vlan_tag))
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules--rewrite"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--copied_rules--rule_set--source_rules--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--copied_rules--rule_set--source_rules--tags"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--copied_rules--rule_set--source_rules--vlan_tag"></a>
### Nested Schema for `copied_rules.rule_set.source_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)

