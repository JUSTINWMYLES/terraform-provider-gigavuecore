---
page_title: "gigavuecore_delete_alert_policies Action - gigavuecore"
subcategory: ""
description: |-
  Delete all or given list of alert policies
---

# gigavuecore_delete_alert_policies Action

Delete all or given list of alert policies

## Example Usage

```terraform
action "gigavuecore_delete_alert_policies" "example" {
  config {
    drop_all = true
    policy_names = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `drop_all` (Bool, optional) - If list of policies are not provided this is required
* `policy_names` (List(String), required) - List of alert policies that needs to be deleted
