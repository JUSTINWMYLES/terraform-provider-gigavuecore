---
page_title: "gigavuecore_load_all_icap_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All ICAP Profiles
---

# gigavuecore_load_all_icap_profiles Data Source

Load All ICAP Profiles

## Example Usage

```terraform
data "gigavuecore_load_all_icap_profiles" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
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

* `alias` (String) - Icap Alias
* `cluster_id` (String) - id of the defining cluster
* `exceed_action` (String) - Icap Profile action incase of Http request buffer exceeded
* `http_req_buf` (Number) - Icap Profile Http request buffer in KB
* `inactivity_timeout` (Number) - Icap Inactivity timeout in minutes
* `preview` (Number) - Icap Preview bytes in KB
* `resp_mod` (String) - Icap Response Modification Enable\|Disable
* `resp_timeout` (Number) - Icap Server Response Timeout value in seconds
* `resp_timeout_action` (String) - Response Timeout action Drop\|Bypass
* `server_group` (String) - Icap Server Group Alias
* `src_max_l4_port` (Number) - Icap Service Source l4 port maximum
* `src_min_l4_port` (Number) - Icap Service Source l4 port minimum

