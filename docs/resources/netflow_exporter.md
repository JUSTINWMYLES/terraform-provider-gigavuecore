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
  alias = null
  cluster_id = null
  description = null
  destination = {}
  dscp = null
  filter = {}
  format = null
  nf_version = null
  snmp = {}
  template_refresh = null
  transport = {}
  ttl = null
  tunneled_port = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, optional) - id of the defining cluster
* `description` (String, optional)
* `destination` (Object({address, ip_ver}), required) - Netflow Exporter Destination definition
  * `address` (String, required)
  * `ip_ver` (String, optional)
* `dscp` (Number, optional)
* `filter` (Object({rules}), optional)
  * `rules` (List(Object({pass_rules})), required)
* `format` (String, optional) - export formats 'v9', 'v5' and 'ipfix' applicable only with 'netflow' and '23' is applicable only if format is 'cef'
* `nf_version` (String, optional) - export format versions. '23' is applicable only with format 'cef' and 'v5', 'v9' and 'ipfix' is valid only with format 'netflow'
* `snmp` (Object({enabled}), optional)
  * `enabled` (Bool, optional) - Allow SNMP packets on the tunnel port
* `template_refresh` (Number, optional)
* `transport` (Object({port, protocol}), required) - Netflow Exporter Transport definition
  * `port` (Number, required) - default: 514, when format is CEF and version 23
  * `protocol` (String, optional)
* `ttl` (Number, optional)
* `tunneled_port` (String, optional) - alias of a Tunneled Port to use with this Exporter. Replaced by /ip/interfaces in GigaVUE-OS 5.5.00

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `description` (String, computed)
* `dscp` (Number, computed)
* `filter` (Object({rules}), computed)
  * `rules` (List(Object({pass_rules})), required)
* `format` (String, computed) - export formats 'v9', 'v5' and 'ipfix' applicable only with 'netflow' and '23' is applicable only if format is 'cef'
* `id` (String, computed)
* `nf_version` (String, computed) - export format versions. '23' is applicable only with format 'cef' and 'v5', 'v9' and 'ipfix' is valid only with format 'netflow'
* `snmp` (Object({enabled}), computed)
  * `enabled` (Bool, optional) - Allow SNMP packets on the tunnel port
* `template_refresh` (Number, computed)
* `ttl` (Number, computed)
* `tunneled_port` (String, computed) - alias of a Tunneled Port to use with this Exporter. Replaced by /ip/interfaces in GigaVUE-OS 5.5.00

