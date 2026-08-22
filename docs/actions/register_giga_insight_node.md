---
page_title: "gigavuecore_register_giga_insight_node Action - gigavuecore"
subcategory: ""
description: |-
  Register a GigaInsight Node with FM
---

# gigavuecore_register_giga_insight_node Action

Register a GigaInsight Node with FM

## Example Usage

```terraform
action "gigavuecore_register_giga_insight_node" "example" {
  config {
    ipv4_address = "example"
    ipv6_address = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ipv4_address` (String, required)
* `ipv6_address` (String, required)
