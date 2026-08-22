---
page_title: "gigavuecore_sub_flow Resource - gigavuecore"
subcategory: ""
description: |-
  Get a sub-flow within a flow
---

# gigavuecore_sub_flow Resource

Get a sub-flow within a flow

## Example Usage

```terraform
resource "gigavuecore_sub_flow" "example" {
  flow_rules = {}
  flow_sample5_g_overlap_rules = {}
  flow_sample5_g_rules = {}
  flow_sample_diameter_rules = {}
  flow_sample_overlap_rules = {}
  flow_sample_rules = {}
  flow_sample_sip_rules = {}
  flow_whitelist5_g_overlap_rules = {}
  flow_whitelist5_g_rules = {}
  flow_whitelist_overlap_rules = {}
  flow_whitelist_rules = {}
  gs_rules = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `flow_rules` (Object({drop_rules, pass_rules}), optional) - Map Flow Rules Container. Private class
  * `drop_rules` (Set(Object({gtp, rule_id})), optional)
  * `pass_rules` (Set(Object({gtp, rule_id})), optional)
* `flow_sample5_g_overlap_rules` (Object({pass_rules}), optional) - Map Flow Sample 5g Overlap Rules Container. Private class
  * `pass_rules` (Set(Object({comment, flow5_g, percentage, rule_id})), optional)
* `flow_sample5_g_rules` (Object({pass_rules}), optional) - Map Flow Sample 5g Rules Container. Private class
  * `pass_rules` (Set(Object({comment, flow5_g, percentage, priority, rule_id})), optional)
* `flow_sample_diameter_rules` (Object({pass_rules}), optional) - Map Flow Sample Diameter Rules Container. Private class
  * `pass_rules` (Set(Object({diameter, interface, percentage, rule_id})), optional)
* `flow_sample_overlap_rules` (Object({pass_rules}), optional) - Map Flow Sample Overlap Rules Container. Private class
  * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), optional)
* `flow_sample_rules` (Object({pass_rules}), optional) - Map Flow Sample Rules Container. Private class
  * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), optional)
* `flow_sample_sip_rules` (Object({pass_rules}), optional) - Map Flow Sample Sip Rules Container. Private class
  * `pass_rules` (Set(Object({percentage, rule_id, sip})), optional)
* `flow_whitelist5_g_overlap_rules` (Object({dnn, type}), optional) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class
  * `dnn` (String, optional) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
  * `type` (String, optional) - Set 5G WL-DB lookup type
* `flow_whitelist5_g_rules` (Object({dnn, type, whitelist_databases}), optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
  * `dnn` (String, optional) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
  * `type` (String, optional) - Set 5G WL-DB lookup type
  * `whitelist_databases` (List(String), optional) - Attach whitelist databases to the map
* `flow_whitelist_overlap_rules` (Object({pass_rules}), optional) - Map Flow Whitelist Overlap Rules Container. Private class
  * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), optional)
* `flow_whitelist_rules` (Object({pass_rules}), optional) - Map Flow Whitelist Rules Container. Private class
  * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), optional)
* `gs_rules` (Object({drop_rules, pass_rules}), optional) - Map GigaSMART Rules Container. Private class
  * `drop_rules` (Set(Object({comment, matches, rule_id})), optional)
  * `pass_rules` (Set(Object({comment, matches, rule_id})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `flow_rules` (Object({drop_rules, pass_rules}), computed) - Map Flow Rules Container. Private class
  * `drop_rules` (Set(Object({gtp, rule_id})), optional)
  * `pass_rules` (Set(Object({gtp, rule_id})), optional)
* `flow_sample5_g_overlap_rules` (Object({pass_rules}), computed) - Map Flow Sample 5g Overlap Rules Container. Private class
  * `pass_rules` (Set(Object({comment, flow5_g, percentage, rule_id})), optional)
* `flow_sample5_g_rules` (Object({pass_rules}), computed) - Map Flow Sample 5g Rules Container. Private class
  * `pass_rules` (Set(Object({comment, flow5_g, percentage, priority, rule_id})), optional)
* `flow_sample_diameter_rules` (Object({pass_rules}), computed) - Map Flow Sample Diameter Rules Container. Private class
  * `pass_rules` (Set(Object({diameter, interface, percentage, rule_id})), optional)
* `flow_sample_overlap_rules` (Object({pass_rules}), computed) - Map Flow Sample Overlap Rules Container. Private class
  * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), optional)
* `flow_sample_rules` (Object({pass_rules}), computed) - Map Flow Sample Rules Container. Private class
  * `pass_rules` (Set(Object({comment, gtp, percentage, periodic_recalc, priority, rule_id})), optional)
* `flow_sample_sip_rules` (Object({pass_rules}), computed) - Map Flow Sample Sip Rules Container. Private class
  * `pass_rules` (Set(Object({percentage, rule_id, sip})), optional)
* `flow_whitelist5_g_overlap_rules` (Object({dnn, type}), computed) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class
  * `dnn` (String, optional) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
  * `type` (String, optional) - Set 5G WL-DB lookup type
* `flow_whitelist5_g_rules` (Object({dnn, type, whitelist_databases}), computed) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
  * `dnn` (String, optional) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
  * `type` (String, optional) - Set 5G WL-DB lookup type
  * `whitelist_databases` (List(String), optional) - Attach whitelist databases to the map
* `flow_whitelist_overlap_rules` (Object({pass_rules}), computed) - Map Flow Whitelist Overlap Rules Container. Private class
  * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), optional)
* `flow_whitelist_rules` (Object({pass_rules}), computed) - Map Flow Whitelist Rules Container. Private class
  * `pass_rules` (Set(Object({flow5_g, gtp, rule_id, sip})), optional)
* `gs_rules` (Object({drop_rules, pass_rules}), computed) - Map GigaSMART Rules Container. Private class
  * `drop_rules` (Set(Object({comment, matches, rule_id})), optional)
  * `pass_rules` (Set(Object({comment, matches, rule_id})), optional)
* `id` (String, computed)

