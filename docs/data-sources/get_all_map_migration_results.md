---
page_title: "gigavuecore_get_all_map_migration_results Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all map migration results
---

# gigavuecore_get_all_map_migration_results Data Source

Load all map migration results

## Example Usage

```terraform
data "gigavuecore_get_all_map_migration_results" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias_mismatched` (Map of String)
* `aliases` (List of String)
* `avg_time_per_map` (Number)
* `cluster_id` (String)
* `cluster_ids` (List of String)
* `clusters_with_no_maps` (List of String)
* `dry_run` (Boolean)
* `duplicate_maps` (List of String)
* `end_time` (Number)
* `failed_cluster_and_reasons` (Map of String)
* `failed_maps_and_reasons` (Map of String)
* `failure_reason` (String)
* `is_all_maps` (Boolean)
* `map_type` (String)
* `migrated_maps` (List of String)
* `migration_alias` (String)
* `skipped_migration` (Map of String)
* `start_time` (Number)
* `status` (String)
* `total_duplicates` (Number)
* `total_failed` (Number)
* `total_maps` (Number)
* `total_port_alias_mismatches` (Number)
* `total_skipped` (Number)
* `total_success` (Number)
* `total_time_to_complete` (Number)
* `triggered_by` (String)

