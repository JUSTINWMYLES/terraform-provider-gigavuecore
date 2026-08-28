---
page_title: "gigavuecore_set_user_selected_ciphers Action - gigavuecore"
subcategory: ""
description: |-
  Set user-selected ciphers by protocol type
---

# gigavuecore_set_user_selected_ciphers Action

Set user-selected ciphers by protocol type

## Example Usage

```terraform
action "gigavuecore_set_user_selected_ciphers" "example" {
  config {
    source_details = null
    type           = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `source_details` (List of Dynamic, optional)
* `type` (String, required) - Cipher protocol type (for example, TLS, SSH)


