---
page_title: "gigavuecore_clear_policy_status Action - gigavuecore"
subcategory: ""
description: |-
  Clear policy deployment status
---

# gigavuecore_clear_policy_status Action

Clear policy deployment status

## Example Usage

```terraform
action "gigavuecore_clear_policy_status" "example" {
  config {
    name = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `name` (String, required) - policy name


