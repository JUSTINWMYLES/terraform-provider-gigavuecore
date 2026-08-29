---
page_title: "gigavuecore_ntp_server Resource - gigavuecore"
subcategory: ""
description: |-
  Find NtpServer by address
---

# gigavuecore_ntp_server Resource

Find NtpServer by address

## Example Usage

```terraform
resource "gigavuecore_ntp_server" "example" {
  cluster_id  = "example"
  enabled     = true
  key_enabled = true
  key_number  = 0
  preferred   = true
  server      = "example"
  version     = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional)
* `key_enabled` (Boolean, optional)
* `key_number` (Number, optional)
* `preferred` (Boolean, optional)
* `server` (String, required) - ipv4 or ipv6 or hostname
* `version` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `enabled` (Boolean, computed)
* `key_enabled` (Boolean, computed)
* `key_number` (Number, computed)
* `preferred` (Boolean, computed)
* `version` (String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ntp_server.example {server}
```
