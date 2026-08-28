---
page_title: "gigavuecore_get_all_ssl_decryption_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Ssl Decryption Report Summary
---

# gigavuecore_get_all_ssl_decryption_summary Data Source

Load all Ssl Decryption Report Summary

## Example Usage

```terraform
data "gigavuecore_get_all_ssl_decryption_summary" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `gsgroup` (String) - alias of gsgroup
* `session_ids` (Number)
* `ssl30_session` (Number)
* `tickets` (Number)
* `tls10_session` (Number)
* `tls11_session` (Number)
* `tls12_session` (Number)
* `total_session` (Number)

