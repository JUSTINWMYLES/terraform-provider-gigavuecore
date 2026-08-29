---
page_title: "gigavuecore_get_netflow_exporter_filter Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Netflow Exporter Filter
---

# gigavuecore_get_netflow_exporter_filter Data Source

Get Netflow Exporter Filter

## Example Usage

```terraform
data "gigavuecore_get_netflow_exporter_filter" "example" {
  alias      = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--items--pass_rules))

<a id="nestedatt--items--pass_rules"></a>
### Nested Schema for `items.pass_rules`

Read-Only:

* `matches` (List of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)

