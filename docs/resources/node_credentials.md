---
page_title: "gigavuecore_node_credentials Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Device Credentials
---

# gigavuecore_node_credentials Resource

Create a new Device Credentials

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_node_credentials.example {device_address}
```
