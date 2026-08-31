---
page_title: "gigavuecore_configure_permitted_addresses Action - gigavuecore"
subcategory: ""
description: |-
  Configure Permitted Addresses
---

# gigavuecore_configure_permitted_addresses Action

Configure Permitted Addresses

## Example Usage

```terraform
action "gigavuecore_configure_permitted_addresses" "example" {
  config {
    address           = "example"
    address_type      = "domain"
    alias             = "example"
    subscribed_events = [ "example" ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `address` (String, required) - Permitted address
* `address_type` (String, required) - Address type
* `alias` (String, required) - Alias for the address
* `subscribed_events` (List of String, optional) - List of events subscribed by the address


