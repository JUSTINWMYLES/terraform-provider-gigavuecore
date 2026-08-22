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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `map_migration_results` (List(Object({alias_mismatched, aliases, avg_time_per_map, cluster_id, cluster_ids, clusters_with_no_maps, dry_run, duplicate_maps, end_time, failed_cluster_and_reasons, failed_maps_and_reasons, failure_reason, is_all_maps, map_type, migrated_maps, migration_alias, skipped_migration, start_time, status, total_duplicates, total_failed, total_maps, total_port_alias_mismatches, total_skipped, total_success, total_time_to_complete, triggered_by})), computed)

