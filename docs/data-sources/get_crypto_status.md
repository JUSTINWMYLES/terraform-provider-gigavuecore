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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `enabled` (Boolean, computed) - Crypto status of FM


