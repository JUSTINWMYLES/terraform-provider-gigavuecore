---
page_title: "gigavuecore_load_all_metadata_exporters Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Metadata Exporters
---

# gigavuecore_load_all_metadata_exporters Data Source

Load All Metadata Exporters

## Example Usage

```terraform
data "gigavuecore_load_all_metadata_exporters" "example" {
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--items--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--items--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--items--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--items--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--items--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--items--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--items--source))
* `type` (String)
<a id="nestedatt--items--cef"></a>
### Nested Schema for `items.cef`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--items--destination"></a>
### Nested Schema for `items.destination`

Read-Only:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--items--mobility_sam"></a>
### Nested Schema for `items.mobility_sam`

Read-Only:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--items--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--items--mobility_sam--event_enable"></a>
### Nested Schema for `items.mobility_sam.event_enable`

Read-Only:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--items--monitor"></a>
### Nested Schema for `items.monitor`

Read-Only:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--items--netflow"></a>
### Nested Schema for `items.netflow`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--items--snmp"></a>
### Nested Schema for `items.snmp`

Read-Only:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--items--source"></a>
### Nested Schema for `items.source`

Read-Only:

* `ip_interface` (String)

