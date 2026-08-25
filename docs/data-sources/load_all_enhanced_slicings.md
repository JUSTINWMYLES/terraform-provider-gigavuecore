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
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `enhanced_slicings` (Attributes List, computed) (see [below for nested schema](#nestedatt--enhanced_slicings))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--enhanced_slicings"></a>
### Nested Schema for `enhanced_slicings`

Read-Only:

* `alias` (String) - apps enhanced slicing alias
* `hash_field` (String) - Hash field for session
* `max_sessions` (Number) - Maximum number of session entries (in millions). Used only for flow-session option
* `protocol_fields` (Attributes List) (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields))
<a id="nestedatt--enhanced_slicings--protocol_fields"></a>
### Nested Schema for `enhanced_slicings.protocol_fields`

Read-Only:

* `gtp` (Attributes) (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--gtp))
* `gtp_u` (Attributes) (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--gtp_u))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--ip))
* `none` (Attributes) (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--none))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--transport))
<a id="nestedatt--enhanced_slicings--protocol_fields--gtp"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.gtp`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--gtp--flow_session))
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
<a id="nestedatt--enhanced_slicings--protocol_fields--gtp--flow_session"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.gtp.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--enhanced_slicings--protocol_fields--gtp_u"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.gtp_u`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--gtp_u--flow_session))
* `l4_port` (Number) - Layer 4 Port Number
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
<a id="nestedatt--enhanced_slicings--protocol_fields--gtp_u--flow_session"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.gtp_u.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--enhanced_slicings--protocol_fields--ip"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.ip`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--ip--flow_session))
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
* `value` (String)
<a id="nestedatt--enhanced_slicings--protocol_fields--ip--flow_session"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.ip.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--enhanced_slicings--protocol_fields--none"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.none`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--none--flow_session))
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
<a id="nestedatt--enhanced_slicings--protocol_fields--none--flow_session"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.none.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port
<a id="nestedatt--enhanced_slicings--protocol_fields--transport"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.transport`

Read-Only:

* `flow_session` (Attributes) - Always use 4-tuple (IP, L4 port) to identify a flow. Used only if user wants to start slicing after some packet count from the beginning of a session flow (see [below for nested schema](#nestedatt--enhanced_slicings--protocol_fields--transport--flow_session))
* `l4_port` (Number) - Layer 4 Port Number
* `offset` (Number) - Number of bytes after 'protocol' header. 64-9000 if 'protocol' is none, 0-9000 for all protocols
* `protocol` (String)
* `value` (String)
<a id="nestedatt--enhanced_slicings--protocol_fields--transport--flow_session"></a>
### Nested Schema for `enhanced_slicings.protocol_fields.transport.flow_session`

Read-Only:

* `action` (String) - Slice or drop packets after skip count
* `skip_pkt_count` (Number) - Start slicing after 'packet-count' is reached. Only valid with 'flowSession'
* `timeout` (Number) - Session inactivity timeout in seconds
* `value` (String) - outer: first occurrence of IP and L4 port ; inner: second occurrence of IP and L4 port

