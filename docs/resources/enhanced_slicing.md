---
page_title: "gigavuecore_enhanced_slicing Resource - gigavuecore"
subcategory: ""
description: |-
  new in version H 5.7
---

# gigavuecore_enhanced_slicing Resource

new in version H 5.7

## Example Usage

```terraform
resource "gigavuecore_enhanced_slicing" "example" {
  alias        = "example"
  hash_field   = "5tuple"
  max_sessions = 4
  protocol_fields = [{
    gtp = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      offset   = 0
      protocol = "gtp"
    }
    gtp_u = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      l4_port  = 1
      offset   = 0
      protocol = "gtpu-tcp"
    }
    ip = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      offset   = 0
      protocol = "ip"
      value    = "inner"
    }
    none = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      offset = 0
    }
    transport = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      l4_port  = 1
      offset   = 0
      protocol = "tcp"
      value    = "inner"
    }
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - apps enhanced slicing alias
* `hash_field` (String, optional) - Hash field for session
* `max_sessions` (Number, optional) - Maximum number of session entries (in millions). Used only for flow-session option
* `protocol_fields` (Attributes List, required) (see [below for nested schema](#nestedatt--protocol_fields))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--protocol_fields"></a>
### Nested Schema for `protocol_fields`

Optional:

* `gtp` (Attributes) (see [below for nested schema](#nestedatt--protocol_fields--gtp))
* `gtp_u` (Attributes) (see [below for nested schema](#nestedatt--protocol_fields--gtp_u))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--protocol_fields--ip))
* `none` (Attributes) (see [below for nested schema](#nestedatt--protocol_fields--none))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--protocol_fields--transport))

<a id="nestedatt--protocol_fields--gtp"></a>
### Nested Schema for `protocol_fields.gtp`

Required:

* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)

Optional:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--protocol_fields--gtp--flow_session))

<a id="nestedatt--protocol_fields--gtp--flow_session"></a>
### Nested Schema for `protocol_fields.gtp.flow_session`

Required:

* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'

Optional:

* `action` (String) - Slice or drop packets after skip count
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port

<a id="nestedatt--protocol_fields--gtp_u"></a>
### Nested Schema for `protocol_fields.gtp_u`

Required:

* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)

Optional:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--protocol_fields--gtp_u--flow_session))
* `l4_port` (Number) - Layer 4 Port Number

<a id="nestedatt--protocol_fields--gtp_u--flow_session"></a>
### Nested Schema for `protocol_fields.gtp_u.flow_session`

Required:

* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'

Optional:

* `action` (String) - Slice or drop packets after skip count
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port

<a id="nestedatt--protocol_fields--ip"></a>
### Nested Schema for `protocol_fields.ip`

Required:

* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
* `value` (String)

Optional:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--protocol_fields--ip--flow_session))

<a id="nestedatt--protocol_fields--ip--flow_session"></a>
### Nested Schema for `protocol_fields.ip.flow_session`

Required:

* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'

Optional:

* `action` (String) - Slice or drop packets after skip count
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port

<a id="nestedatt--protocol_fields--none"></a>
### Nested Schema for `protocol_fields.none`

Required:

* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols

Optional:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--protocol_fields--none--flow_session))

<a id="nestedatt--protocol_fields--none--flow_session"></a>
### Nested Schema for `protocol_fields.none.flow_session`

Required:

* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'

Optional:

* `action` (String) - Slice or drop packets after skip count
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port

<a id="nestedatt--protocol_fields--transport"></a>
### Nested Schema for `protocol_fields.transport`

Required:

* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
* `value` (String)

Optional:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--protocol_fields--transport--flow_session))
* `l4_port` (Number) - Layer 4 Port Number

<a id="nestedatt--protocol_fields--transport--flow_session"></a>
### Nested Schema for `protocol_fields.transport.flow_session`

Required:

* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'

Optional:

* `action` (String) - Slice or drop packets after skip count
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
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
terraform import gigavuecore_enhanced_slicing.example {alias}
```
