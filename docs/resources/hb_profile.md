---
page_title: "gigavuecore_hb_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Find Heartbeat Profile by alias
---

# gigavuecore_hb_profile Resource

Find Heartbeat Profile by alias

## Example Usage

```terraform
resource "gigavuecore_hb_profile" "example" {
  alias                   = "example"
  cluster_id              = "example"
  custom_packet           = "example"
  custom_packet_alias     = "example"
  custom_packet_file_name = "example"
  direction               = "aToB"
  packet_format           = "arp"
  period                  = 30
  recovery_time           = 5
  retries                 = 0
  timeout                 = 20
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Heartbeat Profile alias. Unique within a cluster
* `cluster_id` (String, required) - Target Cluster ID
* `custom_packet` (String, optional) - Base64-encoded custom pcap packet. If omitted, a standard ICMP ARP packet will be used as a heartbeat packet. Custom heartbeat packets can be ICMP/IPv4 or ICMPv6/IPv6
* `custom_packet_alias` (String, optional) - (deprecated) Alias of referenced 'custom heartbeat packet' entry. Used for 'inline heartbeat profile' creation
* `custom_packet_file_name` (String, optional) - File name of referenced 'custom heartbeat packet' entry
* `direction` (String, optional)
* `packet_format` (String, optional)
* `period` (Number, optional) - number of milliseconds between sending subsequent heartbeat packets
* `recovery_time` (Number, optional) - the minimum number of seconds with successfully received packet to declare that the inline tool is up
* `retries` (Number, optional) - number of consecutive timed-out heartbeat packets at which the system will trigger a failover condition
* `timeout` (Number, optional) - number of milliseconds allowed for a heartbeat packet between sending and receiving

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hb_profile.example {alias}/{cluster_id}
```
