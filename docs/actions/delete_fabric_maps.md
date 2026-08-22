---
page_title: "gigavuecore_delete_fabric_maps Action - gigavuecore"
subcategory: ""
description: |-
  Delete all user-defined fabric maps
---

# gigavuecore_delete_fabric_maps Action

Delete all user-defined fabric maps

## Example Usage

```terraform
action "gigavuecore_delete_fabric_maps" "example" {
  config {
    async = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `async` (String, optional) - Specifies the deletion mode. If set to \`true\`, the deletion will be performed asynchronously, allowing the request to return immediately while the deletion process continues in the background. If set to \`false\`, the deletion will be performed synchronously, meaning the request will wait until the deletion is complete before returning a response.
