---
page_title: "gigavuecore_add_flow_rules Action - gigavuecore"
subcategory: ""
description: |-
  Add rules within a flow
---

# gigavuecore_add_flow_rules Action

Add rules within a flow

## Example Usage

```terraform
action "gigavuecore_add_flow_rules" "example" {
  config {
    alias                           = "example"
    flow_alias                      = "example"
    flow_rules                      = null
    flow_sample5_g_overlap_rules    = null
    flow_sample5_g_rules            = null
    flow_sample_diameter_rules      = null
    flow_sample_overlap_rules       = null
    flow_sample_rules               = null
    flow_sample_sip_rules           = null
    flow_whitelist5_g_overlap_rules = null
    flow_whitelist5_g_rules         = null
    flow_whitelist_overlap_rules    = null
    flow_whitelist_rules            = null
    gs_rules                        = null
    rule_type                       = "example"
    sub_flow_alias                  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic Flows alias or ID
* `flow_alias` (String, required) - Flow alias or ID
* `flow_rules` (Dynamic, optional) - Map Flow Rules Container. Private class
* `flow_sample5_g_overlap_rules` (Dynamic, optional) - Map Flow Sample 5g Overlap Rules Container. Private class
* `flow_sample5_g_rules` (Dynamic, optional) - Map Flow Sample 5g Rules Container. Private class
* `flow_sample_diameter_rules` (Dynamic, optional) - Map Flow Sample Diameter Rules Container. Private class
* `flow_sample_overlap_rules` (Dynamic, optional) - Map Flow Sample Overlap Rules Container. Private class
* `flow_sample_rules` (Dynamic, optional) - Map Flow Sample Rules Container. Private class
* `flow_sample_sip_rules` (Dynamic, optional) - Map Flow Sample Sip Rules Container. Private class
* `flow_whitelist5_g_overlap_rules` (Dynamic, optional) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class
* `flow_whitelist5_g_rules` (Dynamic, optional) - Map Flow Whitelist 5g Rule GTP match Definition. Private class
* `flow_whitelist_overlap_rules` (Dynamic, optional) - Map Flow Whitelist Overlap Rules Container. Private class
* `flow_whitelist_rules` (Dynamic, optional) - Map Flow Whitelist Rules Container. Private class
* `gs_rules` (Dynamic, optional) - Map GigaSMART Rules Container. Private class
* `rule_type` (String, required) - Type of rule
* `sub_flow_alias` (String, required) - Sub-flow alias or ID


