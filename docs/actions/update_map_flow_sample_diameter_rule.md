---
page_title: "gigavuecore_update_map_flow_sample_diameter_rule Action - gigavuecore"
subcategory: ""
description: |-
  update flowSampleDiameterRule of a 'secondLevel/flowSampleDiameter' map
---

# gigavuecore_update_map_flow_sample_diameter_rule Action

update flowSampleDiameterRule of a 'secondLevel/flowSampleDiameter' map

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_map_flow_sample_diameter_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    diameter = {
      user_name = "*"
    }
    interface  = "s6a"
    percentage = 0
    rule_id    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `body_rule_id` (Number, required)
* `diameter` (Attributes, required) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--diameter))
* `interface` (String, required) - interface type
* `percentage` (Number, required)
* `rule_id` (String, required) - id of the rule to update

<a id="nestedatt--diameter"></a>
### Nested Schema for `diameter`

Optional:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix

