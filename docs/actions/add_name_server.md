---
page_title: "gigavuecore_add_name_server Action - gigavuecore"
subcategory: ""
description: |-
  Add name server
---

# gigavuecore_add_name_server Action

Add name server

## Example Usage

```terraform
action "gigavuecore_add_name_server" "example" {
  config {
    interface_name = "example"
    servers = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `interface_name` (String, required) - Interface in which the nameserver resides
* `servers` (List(String), required) - List of name server address
