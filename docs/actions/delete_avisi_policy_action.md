---
page_title: "gigavuecore_delete_avisi_policy_action Action - gigavuecore"
subcategory: ""
description: |-
  Delete Active Visibility Policy Action
---

# gigavuecore_delete_avisi_policy_action Action

Delete Active Visibility Policy Action

## Example Usage

```terraform
action "gigavuecore_delete_avisi_policy_action" "example" {
  config {
    action_id = "example"
    policy_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `action_id` (String, required) - Policy action Id
* `policy_id` (String, required) - Policy Id
