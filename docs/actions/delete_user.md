---
page_title: "gigavuecore_delete_user Action - gigavuecore"
subcategory: ""
description: |-
  Delete a user
---

# gigavuecore_delete_user Action

Delete a user

## Example Usage

```terraform
action "gigavuecore_delete_user" "example" {
  config {
    username = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `username` (String, required) - username of target user
