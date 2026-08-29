---
page_title: "gigavuecore_load_all_enhanced_slicings Data Source - gigavuecore"
subcategory: ""
description: |-
  new in version H 5.7
---

# gigavuecore_load_all_enhanced_slicings Data Source

new in version H 5.7

## Example Usage

```terraform
data "gigavuecore_load_all_enhanced_slicings" "example" {
  page = "example"
  sort = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - apps enhanced slicing alias
* `hash_field` (String) - Hash field for session
* `max_sessions` (Number) - Maximum number of session entries (in millions). Used only for flow-session option
* `protocol_fields` (Attributes List) (see [below for nested schema](#nestedatt--items--protocol_fields))
<a id="nestedatt--items--protocol_fields"></a>
### Nested Schema for `items.protocol_fields`

Read-Only:

* `gtp` (Attributes) (see [below for nested schema](#nestedatt--items--protocol_fields--gtp))
* `gtp_u` (Attributes) (see [below for nested schema](#nestedatt--items--protocol_fields--gtp_u))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--items--protocol_fields--ip))
* `none` (Attributes) (see [below for nested schema](#nestedatt--items--protocol_fields--none))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--items--protocol_fields--transport))
<a id="nestedatt--items--protocol_fields--gtp"></a>
### Nested Schema for `items.protocol_fields.gtp`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--items--protocol_fields--gtp--flow_session))
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
<a id="nestedatt--items--protocol_fields--gtp--flow_session"></a>
### Nested Schema for `items.protocol_fields.gtp.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--items--protocol_fields--gtp_u"></a>
### Nested Schema for `items.protocol_fields.gtp_u`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--items--protocol_fields--gtp_u--flow_session))
* `l4_port` (Number) - Layer 4 Port Number
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
<a id="nestedatt--items--protocol_fields--gtp_u--flow_session"></a>
### Nested Schema for `items.protocol_fields.gtp_u.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--items--protocol_fields--ip"></a>
### Nested Schema for `items.protocol_fields.ip`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--items--protocol_fields--ip--flow_session))
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
* `value` (String)
<a id="nestedatt--items--protocol_fields--ip--flow_session"></a>
### Nested Schema for `items.protocol_fields.ip.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--items--protocol_fields--none"></a>
### Nested Schema for `items.protocol_fields.none`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--items--protocol_fields--none--flow_session))
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
<a id="nestedatt--items--protocol_fields--none--flow_session"></a>
### Nested Schema for `items.protocol_fields.none.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--items--protocol_fields--transport"></a>
### Nested Schema for `items.protocol_fields.transport`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--items--protocol_fields--transport--flow_session))
* `l4_port` (Number) - Layer 4 Port Number
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
* `value` (String)
<a id="nestedatt--items--protocol_fields--transport--flow_session"></a>
### Nested Schema for `items.protocol_fields.transport.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port

