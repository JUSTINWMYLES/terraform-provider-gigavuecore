---
page_title: "gigavuecore_interface Resource - gigavuecore"
subcategory: ""
description: |-
  Load Ip Interface by alias
---

# gigavuecore_interface Resource

Load Ip Interface by alias

## Example Usage

```terraform
resource "gigavuecore_interface" "example" {
  alias = null
  attach = []
  comment = null
  gateway = null
  gs_groups = []
  hw_address = null
  ip_address = null
  ip_mask = null
  ip_type = null
  mtu = null
  netflow_exporters = []
  tags = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - ip interface name
* `attach` (List(String), optional) - network ports ,tool ports or circuit ports
* `comment` (String, optional)
* `gateway` (String, optional) - gateway ipv4 or ipv6 address
* `gs_groups` (List(String), optional) - Gs Groups associated with the IP Interface
* `hw_address` (String, optional)
* `ip_address` (String, optional) - ipv4/ipv6 address
* `ip_mask` (String, optional) - ipAddress netmask required with ipAddress
* `ip_type` (String, optional)
* `mtu` (Number, optional)
* `netflow_exporters` (List(String), optional) - Netflow Exporters associated with the IP Interface
* `tags` (List(Object({tag_key, tag_values})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `attach` (List(String), computed) - network ports ,tool ports or circuit ports
* `comment` (String, computed)
* `gateway` (String, computed) - gateway ipv4 or ipv6 address
* `gs_groups` (List(String), computed) - Gs Groups associated with the IP Interface
* `hw_address` (String, computed)
* `ip_address` (String, computed) - ipv4/ipv6 address
* `ip_mask` (String, computed) - ipAddress netmask required with ipAddress
* `ip_type` (String, computed)
* `mtu` (Number, computed)
* `netflow_exporters` (List(String), computed) - Netflow Exporters associated with the IP Interface
* `tags` (List(Object({tag_key, tag_values})), computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_interface.example {alias}
```
