---
page_title: "gigavuecore_load_all_port_throttles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Port Throttles
---

# gigavuecore_load_all_port_throttles Data Source

Load all Port Throttles

## Example Usage

```terraform
data "gigavuecore_load_all_port_throttles" "example" {
  page = null
  sort = null
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

* `alias` (String) - alias of port throttle
* `ports_throttles` (Attributes List) (see [below for nested schema](#nestedatt--items--ports_throttles))
<a id="nestedatt--items--ports_throttles"></a>
### Nested Schema for `items.ports_throttles`

Read-Only:

* `port` (String) - ports or gigastreams
* `type` (String)
* `value` (Number) - Port throttle value

