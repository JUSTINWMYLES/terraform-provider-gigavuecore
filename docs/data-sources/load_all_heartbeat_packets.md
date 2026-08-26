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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Heartbeat Custom Packet alias. Maps directly (foreign key) into InlineHbProfile alias
* `custom_packet` (String) - Base64-encoded custom pcap packet

