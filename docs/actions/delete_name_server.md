---
page_title: "gigavuecore_delete_name_server Action - gigavuecore"
subcategory: ""
description: |-
  Delete name server by address
---

# gigavuecore_delete_name_server Action

Delete name server by address

## Example Usage

```terraform
action "gigavuecore_delete_name_server" "example" {
  config {
    server_ip = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `server_ip` (String, required) - address of the nameserver to be deleted
