---
page_title: "gigavuecore_add_map_flow_sample5_g_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowSample5gRule to a 'secondLevel/flowSample5g' map
---

# gigavuecore_add_map_flow_sample5_g_rule Action

Add new flowSample5gRule to a 'secondLevel/flowSample5g' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_sample5_g_rule" "example" {
  config {
    alias   = "example"
    comment = "example"
    flow5_g = {
      dnn     = "example"
      gpsi    = "example"
      nci     = "example"
      nsiid   = "example"
      pei     = "example"
      plmn_id = "example"
      supi    = "example"
      tac     = "example"
    }
    percentage = 0
    priority   = 0
    rule_id    = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `comment` (String, optional)
* `flow5_g` (Attributes, required) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--flow5_g))
* `percentage` (Number, required)
* `priority` (Number, optional)
* `rule_id` (Number, required)

<a id="nestedatt--flow5_g"></a>
### Nested Schema for `flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix

