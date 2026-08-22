---
page_title: "gigavuecore_metadata_exporter Resource - gigavuecore"
subcategory: ""
description: |-
  Load Metadata Exporter by alias
---

# gigavuecore_metadata_exporter Resource

Load Metadata Exporter by alias

## Example Usage

```terraform
resource "gigavuecore_metadata_exporter" "example" {
  alias = null
  application_profiles = []
  cef = {}
  description = null
  destination = {}
  max_pkt_size = null
  mobility_sam = {}
  monitor = {}
  netflow = {}
  snmp = {}
  source = {}
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `application_profiles` (List(String), optional) - application profile aliases to attach to the exporter
* `cef` (Object({active_timeout, inactive_timeout}), optional)
  * `active_timeout` (Number, optional) - in seconds
  * `inactive_timeout` (Number, optional) - in seconds
* `description` (String, optional)
* `destination` (Object({dscp, ipv4_address, l4_port_dst, l4_port_src, l4_protocol, ttl}), optional)
  * `dscp` (Number, optional)
  * `ipv4_address` (String, optional) - ipv4 address
  * `l4_port_dst` (Number, optional)
  * `l4_port_src` (Number, optional)
  * `l4_protocol` (String, optional)
  * `ttl` (Number, optional)
* `max_pkt_size` (Number, optional)
* `mobility_sam` (Object({encoding, encoding_format, event_enable, trigger}), optional)
  * `encoding` (String, optional)
  * `encoding_format` (String, optional)
  * `event_enable` (Object({modify, update}), optional)
    * `modify` (Bool, optional)
    * `update` (Bool, optional)
  * `trigger` (String, optional)
* `monitor` (Object({timeout}), optional)
  * `timeout` (Number, optional) - how often to export in seconds
* `netflow` (Object({active_timeout, inactive_timeout, template_refresh, template_type, version}), optional)
  * `active_timeout` (Number, optional) - in seconds
  * `inactive_timeout` (Number, optional) - in seconds
  * `template_refresh` (Number, optional) - template refresh interval in seconds
  * `template_type` (String, optional)
  * `version` (String, optional)
* `snmp` (Object({enabled}), optional)
  * `enabled` (Bool, optional) - snmp reverse lookup enable/disable
* `source` (Object({ip_interface}), optional)
  * `ip_interface` (String, optional)
* `type` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `application_profiles` (List(String), computed) - application profile aliases to attach to the exporter
* `cef` (Object({active_timeout, inactive_timeout}), computed)
  * `active_timeout` (Number, optional) - in seconds
  * `inactive_timeout` (Number, optional) - in seconds
* `description` (String, computed)
* `destination` (Object({dscp, ipv4_address, l4_port_dst, l4_port_src, l4_protocol, ttl}), computed)
  * `dscp` (Number, optional)
  * `ipv4_address` (String, optional) - ipv4 address
  * `l4_port_dst` (Number, optional)
  * `l4_port_src` (Number, optional)
  * `l4_protocol` (String, optional)
  * `ttl` (Number, optional)
* `id` (String, computed)
* `max_pkt_size` (Number, computed)
* `mobility_sam` (Object({encoding, encoding_format, event_enable, trigger}), computed)
  * `encoding` (String, optional)
  * `encoding_format` (String, optional)
  * `event_enable` (Object({modify, update}), optional)
    * `modify` (Bool, optional)
    * `update` (Bool, optional)
  * `trigger` (String, optional)
* `monitor` (Object({timeout}), computed)
  * `timeout` (Number, optional) - how often to export in seconds
* `netflow` (Object({active_timeout, inactive_timeout, template_refresh, template_type, version}), computed)
  * `active_timeout` (Number, optional) - in seconds
  * `inactive_timeout` (Number, optional) - in seconds
  * `template_refresh` (Number, optional) - template refresh interval in seconds
  * `template_type` (String, optional)
  * `version` (String, optional)
* `snmp` (Object({enabled}), computed)
  * `enabled` (Bool, optional) - snmp reverse lookup enable/disable
* `source` (Object({ip_interface}), computed)
  * `ip_interface` (String, optional)
* `type` (String, computed)

