---
page_title: "gigavuecore_get_ssl_trust_store_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Get the inline SSL trust-store file
---

# gigavuecore_get_ssl_trust_store_file Data Source

Get the inline SSL trust-store file

## Example Usage

```terraform
data "gigavuecore_get_ssl_trust_store_file" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `certificate` (String, computed)

