---
page_title: "gigavuecore_get_map_group Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve map group for a given policy and source rule
---

# gigavuecore_get_map_group Data Source

Retrieve map group for a given policy and source rule

## Example Usage

```terraform
data "gigavuecore_get_map_group" "example" {
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
* `map_group_config` (Object({cluster_name, comment, components, map_group_alias}), computed)
  * `cluster_name` (String, computed)
  * `comment` (String, computed)
  * `components` (List(Object({alias, health_state, health_state_reasons, source_alias, traffic_health_state, traffic_health_state_reasons, type})), computed)
  * `map_group_alias` (String, computed)
* `source_alias` (String, computed)

