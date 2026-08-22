---
page_title: "gigavuecore_create_bindings Action - gigavuecore"
subcategory: ""
description: |-
  Create an array of bindings to various license targets (chassis or card type, but not both)
---

# gigavuecore_create_bindings Action

Create an array of bindings to various license targets (chassis or card type, but not both)

## Example Usage

```terraform
action "gigavuecore_create_bindings" "example" {
  config {
    activation_id = "example"
    bindings = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, required) - FM License Key bindings to create
* `bindings` (List(Dynamic), optional)
