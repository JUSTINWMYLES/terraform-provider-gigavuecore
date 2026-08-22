---
page_title: "gigavuecore_get_map_migration_results Data Source - gigavuecore"
subcategory: ""
description: |-
  Query map migration results with filters
---

# gigavuecore_get_map_migration_results Data Source

Query map migration results with filters

## Example Usage

```terraform
data "gigavuecore_get_map_migration_results" "example" {
  cluster_id = null
  is_all_maps = null
  is_dry_run = null
  map_alias = null
  map_type = null
  migration_alias = null
  page = null
  sort = null
  status = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID
* `is_all_maps` (Bool, optional) - Include all maps
* `is_dry_run` (Bool, optional) - Filter by dry run status
* `map_alias` (String, optional) - Map alias
* `map_type` (String, optional) - Map type
* `migration_alias` (String, optional) - Migration alias
* `page` (String, optional) - Pagination request string (e.g., "(1:30)")
* `sort` (String, optional) - Sort request string (e.g., "(startTime:DESC)")
* `status` (String, optional) - Migration status

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `map_migration_results` (List(Object({alias_mismatched, aliases, avg_time_per_map, cluster_id, cluster_ids, clusters_with_no_maps, dry_run, duplicate_maps, end_time, failed_cluster_and_reasons, failed_maps_and_reasons, failure_reason, is_all_maps, map_type, migrated_maps, migration_alias, skipped_migration, start_time, status, total_duplicates, total_failed, total_maps, total_port_alias_mismatches, total_skipped, total_success, total_time_to_complete, triggered_by})), computed)

