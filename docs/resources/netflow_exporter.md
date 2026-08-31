---
page_title: "gigavuecore_netflow_exporter Resource - gigavuecore"
subcategory: ""
description: |-
  Find Netflow Exporter by alias
---

# gigavuecore_netflow_exporter Resource

Find Netflow Exporter by alias

## Example Usage

```terraform
resource "gigavuecore_netflow_exporter" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  destination = {
    address = "example"
    ip_ver  = "v4"
  }
  dscp = 0
  filter = {
    rules = [{
      pass_rules = [{
        matches = [ "example" ]
        rule_id = 1
      }]
    }]
  }
  format     = "netflow"
  nf_version = "v5"
  snmp = {
    enabled = true
  }
  template_refresh = 1
  transport = {
    port     = 0
    protocol = "udp"
  }
  ttl           = 1
  tunneled_port = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - id of the defining cluster
* `description` (String, optional)
* `destination` (Attributes, required) - Netflow Exporter Destination definition (see [below for nested schema](#nestedatt--destination))
* `dscp` (Number, optional)
* `filter` (Attributes, optional) (see [below for nested schema](#nestedatt--filter))
* `format` (String, optional) - export formats 'v9', 'v5' and 'ipfix' applicable only with 'netflow' and '23' is applicable only if format is 'cef'
* `nf_version` (String, optional) - export format versions. '23' is applicable only with format 'cef' and 'v5', 'v9' and 'ipfix' is valid only with format 'netflow'
* `snmp` (Attributes, optional) (see [below for nested schema](#nestedatt--snmp))
* `template_refresh` (Number, optional)
* `transport` (Attributes, required) - Netflow Exporter Transport definition (see [below for nested schema](#nestedatt--transport))
* `ttl` (Number, optional)
* `tunneled_port` (String, optional) - alias of a Tunneled Port to use with this Exporter. Replaced by /ip/interfaces in GigaVUE-OS 5.5.00

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Required:

* `address` (String)

Optional:

* `ip_ver` (String)

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

* `rules` (Attributes List) (see [below for nested schema](#nestedatt--filter--rules))

<a id="nestedatt--filter--rules"></a>
### Nested Schema for `filter.rules`

Optional:

* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--filter--rules--pass_rules))

<a id="nestedatt--filter--rules--pass_rules"></a>
### Nested Schema for `filter.rules.pass_rules`

Required:

* `rule_id` (Number)

Optional:

* `matches` (List of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique

<a id="nestedatt--snmp"></a>
### Nested Schema for `snmp`

Optional:

* `enabled` (Boolean) - Allow SNMP packets on the tunnel port

<a id="nestedatt--transport"></a>
### Nested Schema for `transport`

Required:

* `port` (Number) - default: 514, when format is CEF and version 23

Optional:

* `protocol` (String)
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_netflow_exporter.example {alias}:{cluster_id}
```
