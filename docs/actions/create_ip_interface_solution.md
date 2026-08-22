---
page_title: "gigavuecore_create_ip_interface_solution Action - gigavuecore"
subcategory: ""
description: |-
  Define/redefine Ip Interface solution configuration
---

# gigavuecore_create_ip_interface_solution Action

Define/redefine Ip Interface solution configuration

## Example Usage

```terraform
action "gigavuecore_create_ip_interface_solution" "example" {
  config {
    ip_interface_configs = "example"
    tags = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ip_interface_configs` (List(Dynamic), required)
* `tags` (List(Dynamic), optional) - RBAC Tags
