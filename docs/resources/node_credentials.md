---
page_title: "gigavuecore_node_credentials Resource - gigavuecore"
subcategory: ""
description: |-
  Find Device Credentials by address
---

# gigavuecore_node_credentials Resource

Find Device Credentials by address

## Example Usage

```terraform
resource "gigavuecore_node_credentials" "example" {
  device_address = "example"
  hostname       = "example"
  http_password  = "example"
  http_username  = "example"
  https_port     = "example"
  snmp_version   = "v2"
}
```

## Schema

### Arguments

The following arguments are supported:

* `device_address` (String, required) - value of '0.0.0.0' represents the default (fallback) device credentials
* `hostname` (String, optional) - device configured hostname
* `http_password` (String, required) - password to use for device login. On reads, '\*\*\*\*\*' is returned
* `http_username` (String, required) - username to use for device login
* `https_port` (String, optional) - httpsPort to use for device communication. By default 443 is used, If changed the same should be given here
* `snmp_version` (String, optional) - SNMP version to use.

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_node_credentials.example {device_address}
```
