---
page_title: "gigavuecore_load_all_redundancy_profile Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Redundancy Profiles
---

# gigavuecore_load_all_redundancy_profile Data Source

Load all Redundancy Profiles

## Example Usage

```terraform
data "gigavuecore_load_all_redundancy_profile" "example" {
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `redundancy_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--redundancy_profiles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--redundancy_profiles"></a>
### Nested Schema for `redundancy_profiles`

Read-Only:

* `alias` (String) - Inline Network alias. Unique within a cluster
* `protection_role` (String) - only applicable for 'protected' inline networks
* `signaling_port` (String) - only applicable for 'protected' inline networks

