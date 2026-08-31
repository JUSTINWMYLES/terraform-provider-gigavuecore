---
page_title: "gigavuecore_add_map_flow_sample_sip_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new flowSampleSipRule to a 'secondLevel/flowSampleSip' map
---

# gigavuecore_add_map_flow_sample_sip_rule Action

Add new flowSampleSipRule to a 'secondLevel/flowSampleSip' map

## Example Usage

```terraform
action "gigavuecore_add_map_flow_sample_sip_rule" "example" {
  config {
    alias      = "example"
    percentage = 0
    rule_id    = 1
    sip = {
      callee_id = "example"
      callee_id_range = {
        max_value = "example"
        value     = "example"
      }
      caller_id = "example"
      caller_id_range = {
        max_value = "example"
        value     = "example"
      }
      id_range = {
        max_value = "example"
        value     = "example"
      }
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `percentage` (Number, required)
* `rule_id` (Number, required)
* `sip` (Attributes, required) (see [below for nested schema](#nestedatt--sip))

<a id="nestedatt--sip"></a>
### Nested Schema for `sip`

Optional:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--sip--id_range))

<a id="nestedatt--sip--callee_id_range"></a>
### Nested Schema for `sip.callee_id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--sip--caller_id_range"></a>
### Nested Schema for `sip.caller_id_range`

Required:

* `max_value` (String)
* `value` (String)

<a id="nestedatt--sip--id_range"></a>
### Nested Schema for `sip.id_range`

Required:

* `max_value` (String)
* `value` (String)

