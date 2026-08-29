---
page_title: "gigavuecore_engine Resource - gigavuecore"
subcategory: ""
description: |-
  Load GigaSMART port interface details
---

# gigavuecore_engine Resource

Load GigaSMART port interface details

## Example Usage

```terraform
resource "gigavuecore_engine" "example" {
  cluster_id           = "example"
  dhcp                 = true
  dns                  = "example"
  eport                = "example"
  gateway              = "example"
  hw_address           = "example"
  interface            = "example"
  ip_address           = "example"
  ip_mask              = "example"
  mtu                  = 0
  proxy_server_profile = "example"
  status               = "example"
  vlan                 = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `dhcp` (Boolean, optional)
* `dns` (String, optional) - Only valid and required when 'dhcp' is false
* `eport` (String, required) - GigaSMART engine port
* `gateway` (String, optional) - Only valid and required when 'dhcp' is false
* `hw_address` (String, optional)
* `interface` (String, optional)
* `ip_address` (String, optional) - Only valid and required when 'dhcp' is false
* `ip_mask` (String, optional) - Only valid and required when 'dhcp' is false
* `mtu` (Number, optional) - Only valid when 'dhcp' is false
* `proxy_server_profile` (String, optional)
* `status` (String, optional)
* `vlan` (Number, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `dhcp` (Boolean, computed)
* `dns` (String, computed) - Only valid and required when 'dhcp' is false
* `gateway` (String, computed) - Only valid and required when 'dhcp' is false
* `hw_address` (String, computed)
* `interface` (String, computed)
* `ip_address` (String, computed) - Only valid and required when 'dhcp' is false
* `ip_mask` (String, computed) - Only valid and required when 'dhcp' is false
* `mtu` (Number, computed) - Only valid when 'dhcp' is false
* `proxy_server_profile` (String, computed)
* `status` (String, computed)
* `vlan` (Number, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_engine.example {eport}
```
