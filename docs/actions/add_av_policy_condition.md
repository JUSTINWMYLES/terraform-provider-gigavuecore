---
page_title: "gigavuecore_add_av_policy_condition Action - gigavuecore"
subcategory: ""
description: |-
  Add Active Visibility Policy Condition
---

# gigavuecore_add_av_policy_condition Action

Add Active Visibility Policy Condition

## Example Usage

```terraform
action "gigavuecore_add_av_policy_condition" "example" {
  config {
    condition = "example"
    params = null
    policy_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `condition` (String, required) - reference to a pre-defined condition template
* `params` (List(Dynamic), optional) - key/value criteria parameters to instantiate the referenced template
* `policy_id` (String, required) - Policy Id
