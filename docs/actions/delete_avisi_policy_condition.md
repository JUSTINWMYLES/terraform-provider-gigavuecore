---
page_title: "gigavuecore_delete_avisi_policy_condition Action - gigavuecore"
subcategory: ""
description: |-
  Delete Active Visibility Policy Condition
---

# gigavuecore_delete_avisi_policy_condition Action

Delete Active Visibility Policy Condition

## Example Usage

```terraform
action "gigavuecore_delete_avisi_policy_condition" "example" {
  config {
    condition_id = "example"
    policy_id    = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `condition_id` (String, required) - Policy Condition Id
* `policy_id` (String, required) - Policy Id


