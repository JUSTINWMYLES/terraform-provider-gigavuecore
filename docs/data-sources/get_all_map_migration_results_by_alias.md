---
page_title: "gigavuecore_get_all_map_migration_results_by_alias Data Source - gigavuecore"
subcategory: ""
description: |-
  Load map migration result by alias
---

# gigavuecore_get_all_map_migration_results_by_alias Data Source

Load map migration result by alias

## Example Usage

```terraform
data "gigavuecore_get_all_map_migration_results_by_alias" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Migration alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias_mismatched` (Map(String), computed)
* `aliases` (List(String), computed)
* `avg_time_per_map` (Number, computed)
* `cluster_id` (String, computed)
* `cluster_ids` (List(String), computed)
* `clusters_with_no_maps` (List(String), computed)
* `dry_run` (Bool, computed)
* `duplicate_maps` (List(String), computed)
* `end_time` (Number, computed)
* `failed_cluster_and_reasons` (Map(String), computed)
* `failed_maps_and_reasons` (Map(String), computed)
* `failure_reason` (String, computed)
* `is_all_maps` (Bool, computed)
* `map_type` (String, computed)
* `migrated_maps` (List(String), computed)
* `migration_alias` (String, computed)
* `skipped_migration` (Map(String), computed)
* `start_time` (Number, computed)
* `status` (String, computed)
* `total_duplicates` (Number, computed)
* `total_failed` (Number, computed)
* `total_maps` (Number, computed)
* `total_port_alias_mismatches` (Number, computed)
* `total_skipped` (Number, computed)
* `total_success` (Number, computed)
* `total_time_to_complete` (Number, computed)
* `triggered_by` (String, computed)

