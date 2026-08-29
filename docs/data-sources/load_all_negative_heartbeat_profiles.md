---
page_title: "gigavuecore_load_all_negative_heartbeat_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Negative Heartbeat Profiles
---

# gigavuecore_load_all_negative_heartbeat_profiles Data Source

Load all Negative Heartbeat Profiles

## Example Usage

```terraform
data "gigavuecore_load_all_negative_heartbeat_profiles" "example" {
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

* `alias` (String) - Negative Heartbeat Profile alias. Unique within a cluster
* `custom_packet` (String) - Base64-encoded custom pcap packet. Custom negative heartbeat packets can be ICMP/IPv4 or ICMPv6/IPv6
* `custom_packet_file_name` (String) - File name of referenced 'custom negative heartbeat packet' entry. Used for 'inline negative heartbeat profile' creation
* `direction` (String)
* `period` (Number) - number of milliseconds between sending subsequent heartbeat packets
* `recovery_time` (Number) - the minimum number of seconds with successfully received packet to declare that the inline tool is up

