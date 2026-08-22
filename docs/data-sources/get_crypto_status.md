---
page_title: "gigavuecore_get_crypto_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get crypto status.
---

# gigavuecore_get_crypto_status Data Source

Get crypto status.

## Example Usage

```terraform
data "gigavuecore_get_crypto_status" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `enabled` (Bool, computed) - Crypto status of FM

