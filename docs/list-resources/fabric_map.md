---
page_title: "gigavuecore_fabric_map List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all user-defined fabric maps
---

# gigavuecore_fabric_map List Resource

Get all user-defined fabric maps

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_fabric_map" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    alias        = "example"
    dst_ports    = "example"
    health_state = "example"
    page         = "example"
    sort         = "example"
    src_ports    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of fabric map to filter by
* `dst_ports` (String, optional) - Destination port of fabric map to filter by
* `health_state` (String, optional) - HealthState of fabric map to filter by. Possible values are red,green,yellow
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC and default sort field is fabric map alias. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `src_ports` (String, optional) - Source port of fabric map to filter by


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


