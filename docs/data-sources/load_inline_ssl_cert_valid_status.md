---
page_title: "gigavuecore_load_inline_ssl_cert_valid_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Inline SSL certificate validation cache status
---

# gigavuecore_load_inline_ssl_cert_valid_status Data Source

Get Inline SSL certificate validation cache status

## Example Usage

```terraform
data "gigavuecore_load_inline_ssl_cert_valid_status" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `enabled` (Boolean, computed)
* `num_entries` (Number, computed)


