---
page_title: "gigavuecore_add_port_filter_rule Action - gigavuecore"
subcategory: ""
description: |-
  Add new filtering rule to a Port
---

# gigavuecore_add_port_filter_rule Action

Add new filtering rule to a Port

## Example Usage

```terraform
action "gigavuecore_add_port_filter_rule" "example" {
  config {
    cluster_id = "example"
    comment    = "example"
    matches    = [ "example" ]
    port_id    = "example"
    rule_id    = 1
    rule_type  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `port_id` (String, required) - id of the target device Port (format: boxId\_slotId\_port, example: 1\_1\_c1)
* `rule_id` (Number, required)
* `rule_type` (String, required) - filter rule type


