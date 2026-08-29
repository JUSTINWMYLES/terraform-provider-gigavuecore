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
  cluster_id = "example"
  port       = "example"
  rules = {
    drop_rules = [{
      comment = "example"
      matches = [ "example" ]
      rule_id = 0
    }]
    pass_rules = [{
      comment = "example"
      matches = [ "example" ]
      rule_id = 0
    }]
  }
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port` (String, required)
* `rules` (Attributes, required) - Port Filter Rules Container. Private class (see [below for nested schema](#nestedatt--rules))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_filter.example {port}/{cluster_id}
```
