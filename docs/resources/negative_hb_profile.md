---
page_title: "gigavuecore_negative_hb_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Find Negative Heartbeat Profile by alias
---

# gigavuecore_negative_hb_profile Resource

Find Negative Heartbeat Profile by alias

## Example Usage

```terraform
resource "gigavuecore_negative_hb_profile" "example" {
  alias                   = null
  cluster_id              = null
  custom_packet           = null
  custom_packet_file_name = null
  direction               = null
  period                  = null
  recovery_time           = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Negative Heartbeat Profile alias. Unique within a cluster
* `cluster_id` (String, required) - Target Cluster ID
* `custom_packet` (String, optional) - Base64-encoded custom pcap packet. Custom negative heartbeat packets can be ICMP/IPv4 or ICMPv6/IPv6
* `custom_packet_file_name` (String, optional) - File name of referenced 'custom negative heartbeat packet' entry. Used for 'inline negative heartbeat profile' creation
* `direction` (String, optional)
* `period` (Number, optional) - number of milliseconds between sending subsequent heartbeat packets
* `recovery_time` (Number, optional) - the minimum number of seconds with successfully received packet to declare that the inline tool is up

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `custom_packet` (String, computed) - Base64-encoded custom pcap packet. Custom negative heartbeat packets can be ICMP/IPv4 or ICMPv6/IPv6
* `custom_packet_file_name` (String, computed) - File name of referenced 'custom negative heartbeat packet' entry. Used for 'inline negative heartbeat profile' creation
* `direction` (String, computed)
* `period` (Number, computed) - number of milliseconds between sending subsequent heartbeat packets
* `recovery_time` (Number, computed) - the minimum number of seconds with successfully received packet to declare that the inline tool is up


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_negative_hb_profile.example {alias}
```
