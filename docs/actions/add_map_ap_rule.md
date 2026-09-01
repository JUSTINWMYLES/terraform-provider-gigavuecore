---
page_title: "gigavuecore_add_map_ap_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new apRule to a 'secondlevel/byRule' map
---

# gigavuecore_add_map_ap_rule Action

Add new apRule to a 'secondlevel/byRule' map

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_add_map_ap_rule" "example" {
  config {
    alias               = "example"
    application_profile = "example"
    rule_id             = 1
    rule_type           = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `application_profile` (String, required) - application profile alias
* `rule_id` (Number, required) - application profile rule Id, should not have same id as gsRules
* `rule_type` (String, required) - map rule type


