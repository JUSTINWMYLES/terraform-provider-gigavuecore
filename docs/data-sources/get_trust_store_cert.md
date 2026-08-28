---
page_title: "gigavuecore_get_trust_store_cert Data Source - gigavuecore"
subcategory: ""
description: |-
  Get the Trust Store Certificate
---

# gigavuecore_get_trust_store_cert Data Source

Get the Trust Store Certificate

## Example Usage

```terraform
data "gigavuecore_get_trust_store_cert" "example" {
  fingerprint = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `fingerprint` (String, required) - fingerprint for the certificate

### Attributes

In addition to all arguments above, the following attributes are exported:

* `certificate` (String, computed)


