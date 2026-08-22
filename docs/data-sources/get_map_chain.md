---
page_title: "gigavuecore_get_map_chain Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve map chain for a given policy and source rule
---

# gigavuecore_get_map_chain Data Source

Retrieve map chain for a given policy and source rule

## Example Usage

```terraform
data "gigavuecore_get_map_chain" "example" {
  policy_alias = null
  source_rules_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `policy_alias` (String, required) - Policy alias
* `source_rules_alias` (String, required) - Source rules alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed)
* `priority_configs` (List(Object({cluster_name, collector_map_alias, map_chain_id, ordered_components, src_ports_as_id})), computed)
* `source_alias` (String, computed)

