---
page_title: "gigavuecore_get_all_traffic_flows Data Source - gigavuecore"
subcategory: ""
description: |-
  List all Traffic Flows with optional filters
---

# gigavuecore_get_all_traffic_flows Data Source

List all Traffic Flows with optional filters

## Example Usage

```terraform
data "gigavuecore_get_all_traffic_flows" "example" {
  alias = null
  config_status = null
  dst_cluster = null
  dst_ports = null
  health_state = null
  page = null
  select = null
  sort = null
  src_cluster = null
  src_ports = null
  summary = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Traffic Flows alias filter
* `config_status` (String, optional) - Configuration status filter
* `dst_cluster` (String, optional) - Destination cluster filter
* `dst_ports` (String, optional) - Destination port(s) filter
* `health_state` (String, optional) - Health state filter
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `select` (String, optional) - Comma-separated list of fields to select
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC and default sort field is fabric map alias. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `src_cluster` (String, optional) - Source cluster filter
* `src_ports` (String, optional) - Source port(s) filter
* `summary` (Bool, optional) - Return summary data
* `type` (String, optional) - Traffic Flows type filter

### Attributes

In addition to all arguments above, the following attributes are exported:

* `value` (Dynamic, computed)

