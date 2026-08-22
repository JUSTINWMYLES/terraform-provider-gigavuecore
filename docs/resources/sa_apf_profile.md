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
  alias = null
  bidi = null
  buffering = {}
  cluster_id = null
  packet_count = null
  session_fields = []
  timeout = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `bidi` (Bool, optional) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Object({buffer_count_before_match, enabled, protocol}), optional) - Session-Aware APF Buffereing settings
  * `buffer_count_before_match` (Number, optional) - Maximum number of packets BSAPF will buffer per session before APF match
  * `enabled` (Bool, required)
  * `protocol` (String, optional) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
* `cluster_id` (String, optional) - id of the defining cluster
* `packet_count` (Number, optional) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
* `session_fields` (Set(Object({pos, type})), required) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20
* `timeout` (Number, optional) - in seconds

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `bidi` (Bool, computed) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Object({buffer_count_before_match, enabled, protocol}), computed) - Session-Aware APF Buffereing settings
  * `buffer_count_before_match` (Number, optional) - Maximum number of packets BSAPF will buffer per session before APF match
  * `enabled` (Bool, required)
  * `protocol` (String, optional) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
* `cluster_id` (String, computed) - id of the defining cluster
* `packet_count` (Number, computed) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
* `timeout` (Number, computed) - in seconds

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_sa_apf_profile.example {alias}
```
