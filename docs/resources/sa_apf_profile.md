---
page_title: "gigavuecore_sa_apf_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Find SA-APF Profile by alias
---

# gigavuecore_sa_apf_profile Resource

Find SA-APF Profile by alias

## Example Usage

```terraform
resource "gigavuecore_sa_apf_profile" "example" {
  alias = "example"
  bidi  = true
  buffering = {
    buffer_count_before_match = 3
    enabled                   = true
    protocol                  = "tcp"
  }
  cluster_id   = "example"
  packet_count = 0
  session_fields = [{
    pos  = 1
    type = "ipv4Addr"
  }]
  timeout = 10
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `bidi` (Boolean, optional) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Attributes, optional) - Session-Aware APF Buffereing settings (see [below for nested schema](#nestedatt--buffering))
* `cluster_id` (String, required) - id of the defining cluster
* `packet_count` (Number, optional) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
* `session_fields` (Attributes Set, required) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20 (see [below for nested schema](#nestedatt--session_fields))
* `timeout` (Number, optional) - in seconds

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--buffering"></a>
### Nested Schema for `buffering`

Required:

* `enabled` (Boolean)

Optional:

* `buffer_count_before_match` (Number) - Maximum number of packets BSAPF will buffer per session before APF match
* `protocol` (String) - When buffering is enabled, changing protocol requires reboot before new value takes effect.

<a id="nestedatt--session_fields"></a>
### Nested Schema for `session_fields`

Required:

* `type` (String)

Optional:

* `pos` (Number) - Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported
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
terraform import gigavuecore_sa_apf_profile.example {alias}/{cluster_id}
```
