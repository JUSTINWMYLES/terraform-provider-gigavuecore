---
page_title: "gigavuecore_update_map_flow_sample_overlap_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowRule of a 'secondLevel/flowSampleOverlap' map
---

# gigavuecore_update_map_flow_sample_overlap_rule Action

update flowRule of a 'secondLevel/flowSampleOverlap' map

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_map_flow_sample_overlap_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    comment      = "example"
    gtp = {
      apn       = "example"
      eci       = "a1b2c3d4"
      imei      = "*"
      imsi      = "*"
      interface = "Gn"
      msisdn    = "*"
      nas_5_qi  = "0"
      nci       = "*"
      plmn_id   = "123.45"
      qci       = 0
      snssai    = "0"
      tac       = "abc1"
      tac_5_g   = "*"
      version   = "any"
    }
    percentage      = 0
    periodic_recalc = true
    priority        = 1
    rule_id         = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `comment` (String, optional)
* `gtp` (Attributes, required) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--gtp))
* `percentage` (Number, required)
* `periodic_recalc` (Boolean, optional) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number, optional) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (String, required) - id of the rule to update

<a id="nestedatt--gtp"></a>
### Nested Schema for `gtp`

Optional:

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

