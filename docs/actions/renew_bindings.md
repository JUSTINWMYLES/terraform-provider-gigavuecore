---
page_title: "gigavuecore_renew_bindings Action - gigavuecore"
subcategory: ""
description: |-
  Renew eligible licenses
---

# gigavuecore_renew_bindings Action

Renew eligible licenses

## Example Usage

```terraform
action "gigavuecore_renew_bindings" "example" {
  config {
    renewee_id = "example"
    renewer_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `renewee_id` (String, required) - Renewee ID
* `renewer_id` (String, required) - Renewer ID
