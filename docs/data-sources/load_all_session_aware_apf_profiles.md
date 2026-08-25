---
page_title: "gigavuecore_load_all_session_aware_apf_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all defined SA-APF Profiles
---

# gigavuecore_load_all_session_aware_apf_profiles Data Source

Load all defined SA-APF Profiles

## Example Usage

```terraform
data "gigavuecore_load_all_session_aware_apf_profiles" "example" {
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
* `sa_apf_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--sa_apf_profiles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--sa_apf_profiles"></a>
### Nested Schema for `sa_apf_profiles`

Read-Only:

* `alias` (String)
* `bidi` (Boolean) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Attributes) - Session-Aware APF Buffereing settings (see [below for nested schema](#nestedatt--sa_apf_profiles--buffering))
* `cluster_id` (String) - id of the defining cluster
* `packet_count` (Number) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
* `session_fields` (Attributes Set) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20 (see [below for nested schema](#nestedatt--sa_apf_profiles--session_fields))
* `timeout` (Number) - in seconds
<a id="nestedatt--sa_apf_profiles--buffering"></a>
### Nested Schema for `sa_apf_profiles.buffering`

Read-Only:

* `buffer_count_before_match` (Number) - Maximum number of packets BSAPF will buffer per session before APF match
* `enabled` (Boolean)
* `protocol` (String) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
<a id="nestedatt--sa_apf_profiles--session_fields"></a>
### Nested Schema for `sa_apf_profiles.session_fields`

Read-Only:

* `pos` (Number) - Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported
* `type` (String)

