---
page_title: "gigavuecore_load_all_inline_network_lags Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Inline NetworkLags
---

# gigavuecore_load_all_inline_network_lags Data Source

Load all Inline NetworkLags

## Example Usage

```terraform
data "gigavuecore_load_all_inline_network_lags" "example" {
  cluster_id = null
  page = null
  sort = null
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

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `inline_network_lags` (List(Object({alias, cdp, comment, forwarding_state, health_state, health_state_reasons, inline_networks, lacp, lfp, physical_bypass, redundancy_control_state, redundancy_profile, traffic_path})), computed)

