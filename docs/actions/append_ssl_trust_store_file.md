---
page_title: "gigavuecore_append_ssl_trust_store_file Action - gigavuecore"
subcategory: ""
description: |-
  Append inline SSL trust-store from local file
---

# gigavuecore_append_ssl_trust_store_file Action

Append inline SSL trust-store from local file

## Example Usage

```terraform
action "gigavuecore_append_ssl_trust_store_file" "example" {
  config {
    file = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `file` (String, required) - file to upload to device
