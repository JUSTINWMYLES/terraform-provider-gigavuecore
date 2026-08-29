---
page_title: "gigavuecore_update_map_flow_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowRule of a 'secondLevel/flowFilter' map
---

# gigavuecore_update_map_flow_rule Action

update flowRule of a 'secondLevel/flowFilter' map

## Example Usage

```terraform
action "gigavuecore_update_map_flow_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 0
    cluster_id   = "example"
    gtp = {
      imei      = "example"
      imsi      = "example"
      interface = "example"
      msisdn    = "example"
      version   = "example"
    }
    rule_id   = "example"
    rule_type = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `cluster_id` (String, required) - Target Cluster ID
* `gtp` (Attributes, required) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--gtp))
* `rule_id` (String, required) - id of the rule to update
* `rule_type` (String, required) - map rule type

<a id="nestedatt--gtp"></a>
### Nested Schema for `gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface

