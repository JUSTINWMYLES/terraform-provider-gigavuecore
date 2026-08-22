---
page_title: "gigavuecore_clear_precryption_policy_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Clear Precryption Policy configuration
---

# gigavuecore_clear_precryption_policy_configuration Action

Clear Precryption Policy configuration

## Example Usage

```terraform
action "gigavuecore_clear_precryption_policy_configuration" "example" {
  config {
    id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `id` (String, required) - policy graph id
