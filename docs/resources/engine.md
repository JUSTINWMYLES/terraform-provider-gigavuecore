---
page_title: "gigavuecore_engine Resource - gigavuecore"
subcategory: ""
description: |-
  Configure GigaSMART management interface address
---

# gigavuecore_engine Resource

Configure GigaSMART management interface address

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

## Example Usage

```terraform
resource "gigavuecore_engine" "example" {
  cluster_id           = "example"
  dhcp                 = true
  dns                  = "example"
  eport                = "example"
  gateway              = "example"
  hw_address           = "example"
  interface            = "eth2"
  ip_address           = "example"
  ip_mask              = "example"
  mtu                  = 68
  proxy_server_profile = "example"
  status               = "up"
  vlan                 = 20
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_engine.example {eport}:{cluster_id}
```
