---
page_title: "gigavuecore_get_all_internal_fabric_maps List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all internally generated fabric maps supporting a specific user-defined fabric map
---

# gigavuecore_get_all_internal_fabric_maps List Resource

Get all internally generated fabric maps supporting a specific user-defined fabric map

## Example Usage

```terraform
list "gigavuecore_get_all_internal_fabric_maps" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    alias = "example"
    page  = "example"
    sort  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the fabric map
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)
* `ifm_alias` (String, computed)


