---
page_title: "gigavuecore_port_filter Resource - gigavuecore"
subcategory: ""
description: |-
  get filter of a Port by port ID
---

# gigavuecore_port_filter Resource

get filter of a Port by port ID

## Example Usage

```terraform
resource "gigavuecore_port_filter" "example" {
  port  = null
  rules = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `port` (String, required)
* `rules` (Attributes, required) - Port Filter Rules Container. Private class (see [below for nested schema](#nestedatt--rules))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `port_id` (String, computed)

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--rules--pass_rules))
<a id="nestedatt--rules--drop_rules"></a>
### Nested Schema for `rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--rules--pass_rules"></a>
### Nested Schema for `rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `comment` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_filter.example {port_id}
```
