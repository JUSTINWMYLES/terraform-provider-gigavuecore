---
page_title: "gigavuecore_load_all_header_strip List Resource - gigavuecore"
subcategory: ""
description: |-
  Load header strip for all boxes
---

# gigavuecore_load_all_header_strip List Resource

Load header strip for all boxes

## Example Usage

```terraform
list "gigavuecore_load_all_header_strip" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
    page       = "example"
    sort       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target cluster ID.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `box_id` (String, computed)


