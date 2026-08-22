---
page_title: "gigavuecore_redefine_gs_group_diameter_s6_a_session Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Diameter s6a session
---

# gigavuecore_redefine_gs_group_diameter_s6_a_session Action

Redefine GS Group's Diameter s6a session

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_diameter_s6_a_session" "example" {
  config {
    alias = "example"
    limit = 1
    timeout = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `limit` (Number, optional) - Number of sessions to allocate for Diameter S6A
* `timeout` (Number, optional) - timeout in seconds used to clean inactive sessions
