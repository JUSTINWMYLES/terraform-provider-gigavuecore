---
page_title: "gigavuecore_pcap_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Create a PCAP  profile
---

# gigavuecore_pcap_profile Resource

Create a PCAP  profile

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

## Example Usage

```terraform
resource "gigavuecore_pcap_profile" "example" {
  alias_list = "example"
  cluster_id = "example"
  pcap_configs = [{
    alias        = "example"
    channel_port = "example"
    direction    = "example"
    packet_limit = 1
    pcap_rules_list = [{
      dscp = "af11"
      dst_mac = {
        address = "example"
        mask    = "example"
      }
      dstipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      ether_type       = "example"
      inner_vlan       = 1
      ip4_frag         = "noFrag"
      ip4_ttl          = 0
      ip_ver           = "v4"
      packet_hit_count = 0
      portdst          = 0
      portsrc          = 0
      protocol         = 0
      rule_id          = 1
      src_mac = {
        address = "example"
        mask    = "example"
      }
      srcipv4_addrandmask = {
        address = "example"
        mask    = "example"
      }
      tcpctl = 0
      vlan   = 1
    }]
    port = "example"
  }]
  port_ids = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias_list` (String, required) - List of alias to delete. Accepts multiple values, Comma separated eg: aliasList=alias1,alias2
* `cluster_id` (String, required) - Target Cluster ID
* `pcap_configs` (Attributes List, optional) (see [below for nested schema](#nestedatt--pcap_configs))
* `port_ids` (String, required) - List of ports to filterBy. Accepts multiple values, Comma separated eg: portIds=1/1/x1,1/1/x2

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed)
* `channel_port` (String, computed) - port for capturing on egress direction
* `direction` (String, computed) - tx, rx, both
* `id` (String, computed)
* `packet_limit` (Number, computed) - packet limit on a port for capturing
* `pcap_rules_list` (Attributes List, computed) (see [below for nested schema](#nestedatt--pcap_rules_list))
* `port` (String, computed) - portID

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--pcap_configs"></a>
### Nested Schema for `pcap_configs`

Required:

* `alias` (String)
* `port` (String) - portID

Optional:

* `channel_port` (String) - port for capturing on egress direction
* `direction` (String) - tx, rx, both
* `packet_limit` (Number) - packet limit on a port for capturing
* `pcap_rules_list` (Attributes List) (see [below for nested schema](#nestedatt--pcap_configs--pcap_rules_list))

<a id="nestedatt--pcap_configs--pcap_rules_list"></a>
### Nested Schema for `pcap_configs.pcap_rules_list`

Optional:

* `dscp` (String) - DiffServ Code Point bits
* `dst_mac` (Attributes) - MAC address and mask (see [below for nested schema](#nestedatt--pcap_configs--pcap_rules_list--dst_mac))
* `dstipv4_addrandmask` (Attributes) - Ipv4 address and mask. Private class (see [below for nested schema](#nestedatt--pcap_configs--pcap_rules_list--dstipv4_addrandmask))
* `ether_type` (String) - Ether Type
* `inner_vlan` (Number) - Configure inner-vlan id. valid value is between 1 to 4094
* `ip4_frag` (String) - IP fragmentation bits, with in range \[0..255\]
* `ip4_ttl` (Number) - time to live value, with in range \[0..255\]
* `ip_ver` (String) - IP version number
* `packet_hit_count` (Number) - count of packets matches the configured rule. Applicable for GET
* `portdst` (Number) - destination port number.Valid value is between 0 to 65535
* `portsrc` (Number) - source port number.Valid value is between 0 to 65535
* `protocol` (Number) - protocol number in a range \[0..255\]. Well-known protocols numbers are: 0-ipv6Hop, 1-icmpIpv4, 2-igmp, 4-ipv4, 6-tcp, 17-udp, 41-ipv6, 46-rsvp, 47-gre, 58-icmpIpv6
* `rule_id` (Number) - Rule-id to track the filter rules
* `src_mac` (Attributes) - MAC address and mask (see [below for nested schema](#nestedatt--pcap_configs--pcap_rules_list--src_mac))
* `srcipv4_addrandmask` (Attributes) - Ipv4 address and mask. Private class (see [below for nested schema](#nestedatt--pcap_configs--pcap_rules_list--srcipv4_addrandmask))
* `tcpctl` (Number) - Configure TCP control bits: URG, SYN, ACK, etc
* `vlan` (Number) - Configure vlan id. valid value is between 1 to 4094

<a id="nestedatt--pcap_configs--pcap_rules_list--dst_mac"></a>
### Nested Schema for `pcap_configs.pcap_rules_list.dst_mac`

Optional:

* `address` (String) - Specifies a single MAC address to match
* `mask` (String) - Specifies a mask of MAC address

<a id="nestedatt--pcap_configs--pcap_rules_list--dstipv4_addrandmask"></a>
### Nested Schema for `pcap_configs.pcap_rules_list.dstipv4_addrandmask`

Optional:

* `address` (String) - Ipv4 address ex: 10.1.1.1
* `mask` (String) - Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29

<a id="nestedatt--pcap_configs--pcap_rules_list--src_mac"></a>
### Nested Schema for `pcap_configs.pcap_rules_list.src_mac`

Optional:

* `address` (String) - Specifies a single MAC address to match
* `mask` (String) - Specifies a mask of MAC address

<a id="nestedatt--pcap_configs--pcap_rules_list--srcipv4_addrandmask"></a>
### Nested Schema for `pcap_configs.pcap_rules_list.srcipv4_addrandmask`

Optional:

* `address` (String) - Ipv4 address ex: 10.1.1.1
* `mask` (String) - Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29

<a id="nestedatt--pcap_rules_list"></a>
### Nested Schema for `pcap_rules_list`

Read-Only:

* `dscp` (String) - DiffServ Code Point bits
* `dst_mac` (Attributes) - MAC address and mask (see [below for nested schema](#nestedatt--pcap_rules_list--dst_mac))
* `dstipv4_addrandmask` (Attributes) - Ipv4 address and mask. Private class (see [below for nested schema](#nestedatt--pcap_rules_list--dstipv4_addrandmask))
* `ether_type` (String) - Ether Type
* `inner_vlan` (Number) - Configure inner-vlan id. valid value is between 1 to 4094
* `ip4_frag` (String) - IP fragmentation bits, with in range \[0..255\]
* `ip4_ttl` (Number) - time to live value, with in range \[0..255\]
* `ip_ver` (String) - IP version number
* `packet_hit_count` (Number) - count of packets matches the configured rule. Applicable for GET
* `portdst` (Number) - destination port number.Valid value is between 0 to 65535
* `portsrc` (Number) - source port number.Valid value is between 0 to 65535
* `protocol` (Number) - protocol number in a range \[0..255\]. Well-known protocols numbers are: 0-ipv6Hop, 1-icmpIpv4, 2-igmp, 4-ipv4, 6-tcp, 17-udp, 41-ipv6, 46-rsvp, 47-gre, 58-icmpIpv6
* `rule_id` (Number) - Rule-id to track the filter rules
* `src_mac` (Attributes) - MAC address and mask (see [below for nested schema](#nestedatt--pcap_rules_list--src_mac))
* `srcipv4_addrandmask` (Attributes) - Ipv4 address and mask. Private class (see [below for nested schema](#nestedatt--pcap_rules_list--srcipv4_addrandmask))
* `tcpctl` (Number) - Configure TCP control bits: URG, SYN, ACK, etc
* `vlan` (Number) - Configure vlan id. valid value is between 1 to 4094

<a id="nestedatt--pcap_rules_list--dst_mac"></a>
### Nested Schema for `pcap_rules_list.dst_mac`

Read-Only:

* `address` (String) - Specifies a single MAC address to match
* `mask` (String) - Specifies a mask of MAC address

<a id="nestedatt--pcap_rules_list--dstipv4_addrandmask"></a>
### Nested Schema for `pcap_rules_list.dstipv4_addrandmask`

Read-Only:

* `address` (String) - Ipv4 address ex: 10.1.1.1
* `mask` (String) - Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29

<a id="nestedatt--pcap_rules_list--src_mac"></a>
### Nested Schema for `pcap_rules_list.src_mac`

Read-Only:

* `address` (String) - Specifies a single MAC address to match
* `mask` (String) - Specifies a mask of MAC address

<a id="nestedatt--pcap_rules_list--srcipv4_addrandmask"></a>
### Nested Schema for `pcap_rules_list.srcipv4_addrandmask`

Read-Only:

* `address` (String) - Ipv4 address ex: 10.1.1.1
* `mask` (String) - Ipv4 mask ex:255.255.255.248, or <mask length>, e.g. /29
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

