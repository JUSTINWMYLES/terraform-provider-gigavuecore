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
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Migration alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias_mismatched` (Map of String, computed)
* `aliases` (List of String, computed)
* `avg_time_per_map` (Number, computed)
* `cluster_id` (String, computed)
* `cluster_ids` (List of String, computed)
* `clusters_with_no_maps` (List of String, computed)
* `dry_run` (Boolean, computed)
* `duplicate_maps` (List of String, computed)
* `end_time` (Number, computed)
* `failed_cluster_and_reasons` (Map of String, computed)
* `failed_maps_and_reasons` (Map of String, computed)
* `failure_reason` (String, computed)
* `is_all_maps` (Boolean, computed)
* `map_type` (String, computed)
* `migrated_maps` (List of String, computed)
* `migration_alias` (String, computed)
* `skipped_migration` (Map of String, computed)
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


