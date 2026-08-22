---
page_title: "gigavuecore_clear_traffic_policy_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Clear Prefiltering Policy configuration
---

# gigavuecore_clear_traffic_policy_configuration Action

Clear Prefiltering Policy configuration

## Example Usage

```terraform
action "gigavuecore_clear_traffic_policy_configuration" "example" {
  config {
    id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `id` (String, required) - policy graph id
