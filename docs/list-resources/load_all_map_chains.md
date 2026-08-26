---
page_title: "gigavuecore_load_all_map_chains List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all map chains
---

# gigavuecore_load_all_map_chains List Resource

Load all map chains

## Example Usage

```terraform
list "gigavuecore_load_all_map_chains" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
    map_alias  = "example"
    page       = "example"
    sort       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `map_alias` (String, optional) - Filters the response to include only the mapChain that includes the map of the given alias. Note that the map can be either cluster map or fabric map.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `id` (String, computed)


