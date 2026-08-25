---
page_title: "gigavuecore_load_all_icap_server_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All ICAP Server Groups
---

# gigavuecore_load_all_icap_server_groups Data Source

Load All ICAP Server Groups

## Example Usage

```terraform
data "gigavuecore_load_all_icap_server_groups" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Icap Server Group Alias
* `cluster_id` (String) - id of the defining cluster
* `comment` (String) - Icap Server Group Comment
* `icap_servers` (List of String) - ICAP Servers list separated by comma','

