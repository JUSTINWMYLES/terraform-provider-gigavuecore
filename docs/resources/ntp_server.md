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
  enabled     = null
  key_enabled = null
  key_number  = null
  preferred   = null
  server      = null
  version     = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `enabled` (Boolean, optional)
* `key_enabled` (Boolean, optional)
* `key_number` (Number, optional)
* `preferred` (Boolean, optional)
* `server` (String, required) - ipv4 or ipv6 or hostname
* `version` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `enabled` (Boolean, computed)
* `id` (String, computed)
* `key_enabled` (Boolean, computed)
* `key_number` (Number, computed)
* `preferred` (Boolean, computed)
* `version` (String, computed)


