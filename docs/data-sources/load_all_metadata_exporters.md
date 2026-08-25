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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `metadata_exporters` (Attributes List, computed) (see [below for nested schema](#nestedatt--metadata_exporters))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--metadata_exporters"></a>
### Nested Schema for `metadata_exporters`

Read-Only:

* `alias` (String)
* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--source))
* `type` (String)
<a id="nestedatt--metadata_exporters--cef"></a>
### Nested Schema for `metadata_exporters.cef`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--metadata_exporters--destination"></a>
### Nested Schema for `metadata_exporters.destination`

Read-Only:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--metadata_exporters--mobility_sam"></a>
### Nested Schema for `metadata_exporters.mobility_sam`

Read-Only:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--metadata_exporters--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--metadata_exporters--mobility_sam--event_enable"></a>
### Nested Schema for `metadata_exporters.mobility_sam.event_enable`

Read-Only:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--metadata_exporters--monitor"></a>
### Nested Schema for `metadata_exporters.monitor`

Read-Only:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--metadata_exporters--netflow"></a>
### Nested Schema for `metadata_exporters.netflow`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--metadata_exporters--snmp"></a>
### Nested Schema for `metadata_exporters.snmp`

Read-Only:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--metadata_exporters--source"></a>
### Nested Schema for `metadata_exporters.source`

Read-Only:

* `ip_interface` (String)

