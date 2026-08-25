---
page_title: "gigavuecore_fm_build_info Action - gigavuecore"
subcategory: ""
description: |-
  Get FM build info
---

# gigavuecore_fm_build_info Action

Get FM build info

## Example Usage

```terraform
action "gigavuecore_fm_build_info" "example" {
  config {
    eligible_node_role    = "example"
    entity_id             = "example"
    host_name             = "example"
    management_ip_address = "example"
    name                  = "example"
    node_roles            = "example"
    nodes                 = "example"
    password              = "example"
    public_ip_address     = "example"
    username              = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `eligible_node_role` (String, required) - Eligible Node Role of the HA group
* `entity_id` (String, required) - Entity ID of the HA group
* `host_name` (String, required) - DNS Name or IP Address of the HA group
* `management_ip_address` (String, required) - Management Ip Address of the HA group
* `name` (String, required) - name of the HA group
* `node_roles` (Dynamic, required)
* `nodes` (Dynamic, required)
* `password` (String, required) - Password of the HA group
* `public_ip_address` (String, required) - Public Ip Address of the HA group
* `username` (String, required) - Username of the HA group


