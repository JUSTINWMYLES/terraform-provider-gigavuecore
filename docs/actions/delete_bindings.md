---
page_title: "gigavuecore_delete_bindings Action - gigavuecore"
subcategory: ""
description: |-
  Delete FM License Bindings
---

# gigavuecore_delete_bindings Action

Delete FM License Bindings

## Example Usage

```terraform
action "gigavuecore_delete_bindings" "example" {
  config {
    activation_id = "example"
    bindings      = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, required) - FM License Key bindings to delete
* `bindings` (List of Dynamic, optional)


