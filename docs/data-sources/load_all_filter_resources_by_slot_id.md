---
page_title: "gigavuecore_load_all_filter_resources_by_slot_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Filter Resources by slot Id
---

# gigavuecore_load_all_filter_resources_by_slot_id Data Source

Load Filter Resources by slot Id

## Example Usage

```terraform
data "gigavuecore_load_all_filter_resources_by_slot_id" "example" {
  cluster_id = "example"
  slot_id    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `slot_id` (String, required) - Device card slot id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `filter_template` (String, computed) - alias of filter template
* `lookup_resource_limit` (Number, computed)
* `lookup_resource_used` (Number, computed)
* `map_rules_limit` (Number, computed)
* `map_rules_used` (Number, computed)
* `qualifiers` (List of String, computed) - in use qualifiers
* `tool_port_filter_limit` (Number, computed)
* `tool_port_filter_used` (Number, computed)


