---
page_title: "gigavuecore_register_giga_insight_node Action - gigavuecore"
subcategory: ""
description: |-
  Register a GigaInsight Node with FM
---

# gigavuecore_register_giga_insight_node Action

Register a GigaInsight Node with FM

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


