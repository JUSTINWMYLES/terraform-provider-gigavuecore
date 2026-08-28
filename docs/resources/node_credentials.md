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
  device_address = null
  hostname       = null
  http_password  = null
  http_username  = null
  https_port     = null
  snmp_version   = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `hostname` (String, computed) - device configured hostname
* `https_port` (String, computed) - httpsPort to use for device communication. By default 443 is used, If changed the same should be given here
* `id` (String, computed)
* `snmp_version` (String, computed) - SNMP version to use.


