---
page_title: "gigavuecore_delete_trust_store_cert Action - gigavuecore"
subcategory: ""
description: |-
  Delete the Trust Store Certificate
---

# gigavuecore_delete_trust_store_cert Action

Delete the Trust Store Certificate

## Example Usage

```terraform
action "gigavuecore_delete_trust_store_cert" "example" {
  config {
    fingerprint = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `fingerprint` (String, required) - fingerprint for the certificate


