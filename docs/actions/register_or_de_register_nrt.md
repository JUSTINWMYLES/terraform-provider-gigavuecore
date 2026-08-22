---
page_title: "gigavuecore_register_or_de_register_nrt Action - gigavuecore"
subcategory: ""
description: |-
  Register or deregister near real-time statistics for a traffic flow
---

# gigavuecore_register_or_de_register_nrt Action

Register or deregister near real-time statistics for a traffic flow

## Example Usage

```terraform
action "gigavuecore_register_or_de_register_nrt" "example" {
  config {
    alias = "example"
    operation = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias
* `operation` (String, optional) - Operation to perform (add or delete)
