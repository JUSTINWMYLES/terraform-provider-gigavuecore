---
page_title: "gigavuecore_load_all_ib_pathways Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Inter-broker Pathways
---

# gigavuecore_load_all_ib_pathways Data Source

Load all Inter-broker Pathways

## Example Usage

```terraform
data "gigavuecore_load_all_ib_pathways" "example" {
  page = "example"
  sort = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Inline resilient inter-broker pathway alias
* `comment` (String)
* `min_ports_up` (Number) - minimum number of ports in the 'up' state needed to declare the ib-pathway to be in the 'up' state.
* `operational_state` (String)
* `ports` (List of String) - list of local network ports with same speed
* `traffic_path` (String)

