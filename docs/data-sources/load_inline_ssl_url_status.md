---
page_title: "gigavuecore_load_inline_ssl_url_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Inline SSL URL cache status
---

# gigavuecore_load_inline_ssl_url_status Data Source

Get Inline SSL URL cache status

## Example Usage

```terraform
data "gigavuecore_load_inline_ssl_url_status" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `db_version` (String, computed)
* `enabled` (Bool, computed)
* `num_entries` (Number, computed)

