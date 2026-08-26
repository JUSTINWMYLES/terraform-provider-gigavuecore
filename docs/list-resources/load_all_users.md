---
page_title: "gigavuecore_load_all_users List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Users
---

# gigavuecore_load_all_users List Resource

Load all Users

## Example Usage

```terraform
list "gigavuecore_load_all_users" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    page     = "example"
    sort     = "example"
    username = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `username` (String, optional) - username of the target user


### Identity Attributes

The following identity attributes are exported for each matching result:

* `username` (String, computed)


