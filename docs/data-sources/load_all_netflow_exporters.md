---
page_title: "gigavuecore_load_all_netflow_exporters Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all defined Netflow Exporters
---

# gigavuecore_load_all_netflow_exporters Data Source

Load all defined Netflow Exporters

## Example Usage

```terraform
data "gigavuecore_load_all_netflow_exporters" "example" {
  cluster_id = null
  page       = null
  sort       = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cluster_id` (String) - id of the defining cluster
* `description` (String)
* `destination` (Attributes) - Netflow Exporter Destination definition (see [below for nested schema](#nestedatt--items--destination))
* `dscp` (Number)
* `filter` (Attributes) (see [below for nested schema](#nestedatt--items--filter))
* `format` (String) - export formats 'v9', 'v5' and 'ipfix' applicable only with 'netflow' and '23' is applicable only if format is 'cef'
* `nf_version` (String) - export format versions. '23' is applicable only with format 'cef' and 'v5', 'v9' and 'ipfix' is valid only with format 'netflow'
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--items--snmp))
* `template_refresh` (Number)
* `transport` (Attributes) - Netflow Exporter Transport definition (see [below for nested schema](#nestedatt--items--transport))
* `ttl` (Number)
* `tunneled_port` (String) - alias of a Tunneled Port to use with this Exporter. Replaced by /ip/interfaces in GigaVUE-OS 5.5.00
<a id="nestedatt--items--destination"></a>
### Nested Schema for `items.destination`

Read-Only:

* `address` (String)
* `ip_ver` (String)
<a id="nestedatt--items--filter"></a>
### Nested Schema for `items.filter`

Read-Only:

* `rules` (Attributes List) (see [below for nested schema](#nestedatt--items--filter--rules))
<a id="nestedatt--items--filter--rules"></a>
### Nested Schema for `items.filter.rules`

Read-Only:

* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--items--filter--rules--pass_rules))
<a id="nestedatt--items--filter--rules--pass_rules"></a>
### Nested Schema for `items.filter.rules.pass_rules`

Read-Only:

* `matches` (List of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--items--snmp"></a>
### Nested Schema for `items.snmp`

Read-Only:

* `enabled` (Boolean) - Allow SNMP packets on the tunnel port
<a id="nestedatt--items--transport"></a>
### Nested Schema for `items.transport`

Read-Only:

* `port` (Number) - default: 514, when format is CEF and version 23
* `protocol` (String)

