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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({gsgroup, session_ids, ssl30_session, tickets, tls10_session, tls11_session, tls12_session, total_session})), computed)

