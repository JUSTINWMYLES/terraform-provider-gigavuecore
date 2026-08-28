---
page_title: "gigavuecore_load_all_ib_pathways List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Inter-broker Pathways
---

# gigavuecore_load_all_ib_pathways List Resource

Load all Inter-broker Pathways

## Example Usage

```terraform
list "gigavuecore_load_all_ib_pathways" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    page = "example"
    sort = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


