---
page_title: "gigavuecore_redefine_metadata_exporter Action - gigavuecore"
subcategory: ""
description: |-
  Redefine a metadata exporter
---

# gigavuecore_redefine_metadata_exporter Action

Redefine a metadata exporter

## Example Usage

```terraform
action "gigavuecore_redefine_metadata_exporter" "example" {
  config {
    alias                = "example"
    application_profiles = [ "example" ]
    body_alias           = "example"
    cef = {
      active_timeout   = 0
      inactive_timeout = 0
    }
    description = "example"
    destination = {
      dscp         = 0
      ipv4_address = "example"
      l4_port_dst  = 0
      l4_port_src  = 0
      l4_protocol  = "example"
      ttl          = 0
    }
    max_pkt_size = 0
    mobility_sam = {
      encoding        = "example"
      encoding_format = "example"
      event_enable = {
        modify = true
        update = true
      }
      trigger = "example"
    }
    monitor = {
      timeout = 0
    }
    netflow = {
      active_timeout   = 0
      inactive_timeout = 0
      template_refresh = 0
      template_type    = "example"
      version          = "example"
    }
    snmp = {
      enabled = true
    }
    source = {
      ip_interface = "example"
    }
    type = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - metadata exporter alias
* `application_profiles` (List of String, optional) - application profile aliases to attach to the exporter
* `body_alias` (String, required)
* `cef` (Attributes, optional) - cef attributes (see [below for nested schema](#nestedatt--cef))
* `description` (String, optional)
* `destination` (Attributes, optional) - destination attributes (see [below for nested schema](#nestedatt--destination))
* `max_pkt_size` (Number, optional)
* `mobility_sam` (Attributes, optional) (see [below for nested schema](#nestedatt--mobility_sam))
* `monitor` (Attributes, optional) - monitor attributes (see [below for nested schema](#nestedatt--monitor))
* `netflow` (Attributes, optional) - netflow attributes (see [below for nested schema](#nestedatt--netflow))
* `snmp` (Attributes, optional) (see [below for nested schema](#nestedatt--snmp))
* `source` (Attributes, optional) - source tunnel port (see [below for nested schema](#nestedatt--source))
* `type` (String, optional)

<a id="nestedatt--cef"></a>
### Nested Schema for `cef`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds

<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Optional:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)

<a id="nestedatt--mobility_sam"></a>
### Nested Schema for `mobility_sam`

Optional:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--mobility_sam--event_enable))
* `trigger` (String)

<a id="nestedatt--mobility_sam--event_enable"></a>
### Nested Schema for `mobility_sam.event_enable`

Optional:

* `modify` (Boolean)
* `update` (Boolean)

<a id="nestedatt--monitor"></a>
### Nested Schema for `monitor`

Optional:

* `timeout` (Number) - how often to export in seconds

<a id="nestedatt--netflow"></a>
### Nested Schema for `netflow`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)

<a id="nestedatt--snmp"></a>
### Nested Schema for `snmp`

Optional:

* `enabled` (Boolean) - snmp reverse lookup enable/disable

<a id="nestedatt--source"></a>
### Nested Schema for `source`

Optional:

* `ip_interface` (String)

