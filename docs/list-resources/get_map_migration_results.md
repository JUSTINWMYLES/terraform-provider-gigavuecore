---
page_title: "gigavuecore_get_map_migration_results List Resource - gigavuecore"
subcategory: ""
description: |-
  Query map migration results with filters
---

# gigavuecore_get_map_migration_results List Resource

Query map migration results with filters

## Example Usage

```terraform
list "gigavuecore_get_map_migration_results" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id      = "example"
    is_all_maps     = true
    is_dry_run      = true
    map_alias       = "example"
    map_type        = "example"
    migration_alias = "example"
    page            = "example"
    sort            = "example"
    status          = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID
* `is_all_maps` (Boolean, optional) - Include all maps
* `is_dry_run` (Boolean, optional) - Filter by dry run status
* `map_alias` (String, optional) - Map alias
* `map_type` (String, optional) - Map type
* `migration_alias` (String, optional) - Migration alias
* `page` (String, optional) - Pagination request string (e.g., "(1:30)")
* `sort` (String, optional) - Sort request string (e.g., "(startTime:DESC)")
* `status` (String, optional) - Migration status


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


