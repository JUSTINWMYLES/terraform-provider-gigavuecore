---
page_title: "gigavuecore_port_filter_rule Resource - gigavuecore"
subcategory: ""
description: |-
  Add new filtering rule to a Port
---

# gigavuecore_port_filter_rule Resource

Add new filtering rule to a Port

## Example Usage

```terraform
resource "gigavuecore_port_filter_rule" "example" {
  cluster_id = "example"
  comment    = "example"
  matches    = ["example"]
  port_id    = "example"
  rule_id    = 1
  rule_type  = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `matches` (Set of Dynamic, required) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `port_id` (String, required) - id of the target device Port (format: boxId\_slotId\_port, example: 1\_1\_c1)
* `rule_id` (Number, required) - id of the rule to update
* `rule_type` (String, required) - filter rule type

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_filter_rule.example {port_id}:{cluster_id}
```
