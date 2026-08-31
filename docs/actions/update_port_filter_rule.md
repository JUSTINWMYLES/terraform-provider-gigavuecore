---
page_title: "gigavuecore_update_port_filter_rule Action - gigavuecore"
subcategory: ""
description: |-
  update filtering rule of a Port
---

# gigavuecore_update_port_filter_rule Action

update filtering rule of a Port

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_port_filter_rule" "example" {
  config {
    body_rule_id = 1
    cluster_id   = "example"
    comment      = "example"
    matches      = ["example"]
    port_id      = "example"
    rule_id      = "example"
    rule_type    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `body_rule_id` (Number, required)
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `port_id` (String, required) - id of the target device Port (format: boxId\_slotId\_port, example: 1\_1\_c1)
* `rule_id` (String, required) - id of the rule to update
* `rule_type` (String, required) - filter rule type


