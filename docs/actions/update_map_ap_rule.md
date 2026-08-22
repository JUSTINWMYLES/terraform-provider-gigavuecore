---
page_title: "gigavuecore_update_map_ap_rule Action - gigavuecore"
subcategory: ""
description: |-
  update apRule of a 'secondlevel/byRule' map
---

# gigavuecore_update_map_ap_rule Action

update apRule of a 'secondlevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_update_map_ap_rule" "example" {
  config {
    alias = "example"
    application_profile = "example"
    body_rule_id = 1
    rule_id = "example"
    rule_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `application_profile` (String, required) - application profile alias
* `body_rule_id` (Number, required) - application profile rule Id, should not have same id as gsRules
* `rule_id` (String, required) - id of the rule to update
* `rule_type` (String, required) - map rule type
