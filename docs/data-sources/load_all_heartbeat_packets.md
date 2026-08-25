---
page_title: "gigavuecore_load_all_heartbeat_packets Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Heartbeat Packets (deprecated: use GET /inline/hbProfiles)
---

# gigavuecore_load_all_heartbeat_packets Data Source

Load all Heartbeat Packets (deprecated: use GET /inline/hbProfiles)

## Example Usage

```terraform
data "gigavuecore_load_all_heartbeat_packets" "example" {
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
* `hb_custom_packets` (Attributes List, computed) (see [below for nested schema](#nestedatt--hb_custom_packets))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--hb_custom_packets"></a>
### Nested Schema for `hb_custom_packets`

Read-Only:

* `alias` (String) - Heartbeat Custom Packet alias. Maps directly (foreign key) into InlineHbProfile alias
* `custom_packet` (String) - Base64-encoded custom pcap packet

