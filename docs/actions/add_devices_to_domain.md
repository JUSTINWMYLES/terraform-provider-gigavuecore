---
page_title: "gigavuecore_add_devices_to_domain Action - gigavuecore"
subcategory: ""
description: |-
  add clusters/devices to the FM-managed domain
---

# gigavuecore_add_devices_to_domain Action

add clusters/devices to the FM-managed domain

## Example Usage

```terraform
action "gigavuecore_add_devices_to_domain" "example" {
  config {
    async = true
    node_add_specs = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `async` (Bool, optional) - if provided, the call returns immediately with the \[202 Accepted\] HTTP status code, and the discovery process will complete in the background
* `node_add_specs` (List(Dynamic), required)
