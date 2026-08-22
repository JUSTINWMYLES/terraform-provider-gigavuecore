---
page_title: "gigavuecore_bulk_update_nameservers Action - gigavuecore"
subcategory: ""
description: |-
  update name servers
---

# gigavuecore_bulk_update_nameservers Action

update name servers

## Example Usage

```terraform
action "gigavuecore_bulk_update_nameservers" "example" {
  config {
    body_interface_name = "example"
    interface_name = "example"
    servers = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `body_interface_name` (String, required) - Interface in which the nameserver resides
* `interface_name` (String, required) - interfaceName for which name servers are updated
* `servers` (List(String), required) - List of name server address
