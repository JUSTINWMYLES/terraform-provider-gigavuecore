---
page_title: "gigavuecore_get_fm_crypto_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get system fips settings
---

# gigavuecore_get_fm_crypto_status Data Source

Get system fips settings

## Example Usage

```terraform
data "gigavuecore_get_fm_crypto_status" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `enabled` (Bool, computed) - FIPS mode of FM

