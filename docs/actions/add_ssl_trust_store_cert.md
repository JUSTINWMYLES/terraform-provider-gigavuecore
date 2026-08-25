---
page_title: "gigavuecore_add_ssl_trust_store_cert Action - gigavuecore"
subcategory: ""
description: |-
  Add certificate to the inline SSL trust-store
---

# gigavuecore_add_ssl_trust_store_cert Action

Add certificate to the inline SSL trust-store

## Example Usage

```terraform
action "gigavuecore_add_ssl_trust_store_cert" "example" {
  config {
    file        = "example"
    file_source = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Dynamic, optional) - Remote file source or destination


