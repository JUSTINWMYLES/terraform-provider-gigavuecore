---
page_title: "gigavuecore_negative_hb_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Create Negative Heartbeat Profile
---

# gigavuecore_negative_hb_profile Resource

Create Negative Heartbeat Profile

## Example Usage

```terraform
resource "gigavuecore_negative_hb_profile" "example" {
  alias                   = "example"
  cluster_id              = "example"
  custom_packet           = "example"
  custom_packet_file_name = "example"
  direction               = "aToB"
  period                  = 30
  recovery_time           = 5
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_negative_hb_profile.example {alias}/{cluster_id}
```
