---
page_title: "gigavuecore_load_all_metadata_application_profiles List Resource - gigavuecore"
subcategory: ""
description: |-
  Load All Metadata Application Profiles
---

# gigavuecore_load_all_metadata_application_profiles List Resource

Load All Metadata Application Profiles

## Example Usage

```terraform
list "gigavuecore_load_all_metadata_application_profiles" "example" {
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

* `alias` (String, computed) - application profile alias


