---
page_title: "gigavuecore_load_all_header_strip_aging List Resource - gigavuecore"
subcategory: ""
description: |-
  Load header strip aging for all boxes
---

# gigavuecore_load_all_header_strip_aging List Resource

Load header strip aging for all boxes

## Example Usage

```terraform
list "gigavuecore_load_all_header_strip_aging" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    box_id = "example"
    page   = "example"
    sort   = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `box_id` (String, computed) - device box id. valid range 1 - 64. all is applicable only for post request.


