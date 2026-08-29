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
    buffer_count_before_match = 0
    enabled                   = true
    protocol                  = "example"
  }
  cluster_id   = "example"
  packet_count = 0
  session_fields = [{
    pos  = 0
    type = "example"
  }]
  timeout = 0
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `bidi` (Boolean, computed) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Attributes, computed) - Session-Aware APF Buffereing settings (see [below for nested schema](#nestedatt--buffering))
* `packet_count` (Number, computed) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
* `timeout` (Number, computed) - in seconds

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

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_sa_apf_profile.example {alias}
```
