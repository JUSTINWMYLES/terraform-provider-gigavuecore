---
page_title: "gigavuecore_delete_permitted_address Action - gigavuecore"
subcategory: ""
description: |-
  Delete Permitted Address
---

# gigavuecore_delete_permitted_address Action

Delete Permitted Address

## Example Usage

```terraform
action "gigavuecore_delete_permitted_address" "example" {
  config {
    address = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `address` (String, required) - Permitted Address
