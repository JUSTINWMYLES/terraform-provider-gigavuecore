---
page_title: "gigavuecore_cache Resource - gigavuecore"
subcategory: ""
description: |-
  Load Metadata Cache by Alias
---

# gigavuecore_cache Resource

Load Metadata Cache by Alias

## Example Usage

```terraform
resource "gigavuecore_cache" "example" {
  advance_hash     = true
  alias            = "example"
  description      = "example"
  dpi_inject_limit = 0
  event            = "txnEnd"
  exporters        = ["example"]
  flow_behavior    = "unidir"
  match = {
    datalink = {
      mac_dst = true
      mac_src = true
      vlan    = true
    }
    interface = {
      in_name_width     = 1
      in_physical_width = 2
    }
    ip = {
      version = true
    }
    ipv4 = {
      destination = {
        prefix_min_mask = "example"
      }
      dscp = true
      fragmentation = {
        flags  = true
        offset = true
      }
      header_len = true
      option_map = true
      precedence = true
      protocol   = true
      section = {
        header_size  = 1
        payload_size = 1
      }
      source = {
        prefix_min_mask = "example"
      }
      tos          = true
      total_length = true
      ttl          = true
    }
    ipv6 = {
      destination = {
        prefix_min_mask = "example"
      }
      dscp          = true
      extension_map = true
      flow_label    = true
      fragmentation = {
        flags  = true
        offset = true
      }
      hop_limit = true
      length = {
        header  = true
        payload = true
        total   = true
      }
      next_header = true
      precedence  = true
      section = {
        header_size  = 1
        payload_size = 1
      }
      source = {
        prefix_min_mask = "example"
      }
      traffic_class = true
    }
    transport = {
      dst_port = true
      icmp = {
        ipv4_code = true
        ipv4_type = true
        ipv6_code = true
        ipv6_type = true
      }
      src_port = true
      tcp = {
        ack_number  = true
        dst_port    = true
        flags       = true
        header_len  = true
        seq_number  = true
        src_port    = true
        urgent_ptr  = true
        window_size = true
      }
      udp = {
        dst_port = true
        msg_len  = true
        src_port = true
      }
    }
  }
  multi_collect         = true
  network_profiles      = ["example"]
  observation_domain_id = 0
  sampling = {
    mode                 = "multiRate"
    single_sampling_rate = 10
  }
  size = {
    flows = 1
  }
  timeout = {
    idle = 1
  }
}
```

## Schema

### Arguments

The following arguments are supported:

* `advance_hash` (Boolean, optional) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `alias` (String, required)
* `description` (String, optional)
* `dpi_inject_limit` (Number, optional)
* `event` (String, optional)
* `exporters` (List of String, optional) - alias of metadata exporters to attach this cache
* `flow_behavior` (String, optional) - direction for flow identification
* `match` (Attributes, optional) - match criteria for record generation (see [below for nested schema](#nestedatt--match))
* `multi_collect` (Boolean, optional) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List of String, optional) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number, optional)
* `sampling` (Attributes, optional) (see [below for nested schema](#nestedatt--sampling))
* `size` (Attributes, optional) - size of the flows (see [below for nested schema](#nestedatt--size))
* `timeout` (Attributes, optional) (see [below for nested schema](#nestedatt--timeout))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--match"></a>
### Nested Schema for `match`

Optional:

* `datalink` (Attributes) (see [below for nested schema](#nestedatt--match--datalink))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--match--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--match--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--match--transport))

<a id="nestedatt--match--datalink"></a>
### Nested Schema for `match.datalink`

Optional:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)

<a id="nestedatt--match--interface"></a>
### Nested Schema for `match.interface`

Optional:

* `in_name_width` (Number)
* `in_physical_width` (Number)

<a id="nestedatt--match--ip"></a>
### Nested Schema for `match.ip`

Optional:

* `version` (Boolean)

<a id="nestedatt--match--ipv4"></a>
### Nested Schema for `match.ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--match--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--match--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--match--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--match--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)

<a id="nestedatt--match--ipv4--destination"></a>
### Nested Schema for `match.ipv4.destination`

Optional:

* `prefix_min_mask` (String) - ipv4 destination prefix minimum-mask - netmask or mask length

<a id="nestedatt--match--ipv4--fragmentation"></a>
### Nested Schema for `match.ipv4.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)

<a id="nestedatt--match--ipv4--section"></a>
### Nested Schema for `match.ipv4.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)

<a id="nestedatt--match--ipv4--source"></a>
### Nested Schema for `match.ipv4.source`

Optional:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length

<a id="nestedatt--match--ipv6"></a>
### Nested Schema for `match.ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--match--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--match--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--match--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--match--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--match--ipv6--source))
* `traffic_class` (Boolean)

<a id="nestedatt--match--ipv6--destination"></a>
### Nested Schema for `match.ipv6.destination`

Optional:

* `prefix_min_mask` (String)

<a id="nestedatt--match--ipv6--fragmentation"></a>
### Nested Schema for `match.ipv6.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)

<a id="nestedatt--match--ipv6--length"></a>
### Nested Schema for `match.ipv6.length`

Optional:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)

<a id="nestedatt--match--ipv6--section"></a>
### Nested Schema for `match.ipv6.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)

<a id="nestedatt--match--ipv6--source"></a>
### Nested Schema for `match.ipv6.source`

Optional:

* `prefix_min_mask` (String) - ipv6 source prefix minimum-mask - netmask or mask length

<a id="nestedatt--match--transport"></a>
### Nested Schema for `match.transport`

Optional:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--match--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--match--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--match--transport--udp))

<a id="nestedatt--match--transport--icmp"></a>
### Nested Schema for `match.transport.icmp`

Optional:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)

<a id="nestedatt--match--transport--tcp"></a>
### Nested Schema for `match.transport.tcp`

Optional:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)

<a id="nestedatt--match--transport--udp"></a>
### Nested Schema for `match.transport.udp`

Optional:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)

<a id="nestedatt--sampling"></a>
### Nested Schema for `sampling`

Optional:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)

<a id="nestedatt--size"></a>
### Nested Schema for `size`

Optional:

* `flows` (Number) - size of flows in millions

<a id="nestedatt--timeout"></a>
### Nested Schema for `timeout`

Optional:

* `idle` (Number) - idle timeout in seconds. max value 7days. default 30 min
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
terraform import gigavuecore_cache.example {alias}
```
