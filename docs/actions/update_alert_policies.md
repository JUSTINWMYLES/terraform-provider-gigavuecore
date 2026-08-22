---
page_title: "gigavuecore_update_alert_policies Action - gigavuecore"
subcategory: ""
description: |-
  Update of list of given alert-policies
---

# gigavuecore_update_alert_policies Action

Update of list of given alert-policies

## Example Usage

```terraform
action "gigavuecore_update_alert_policies" "example" {
  config {
    context = null
    policies = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `context` (Dynamic, optional) - Gigamon query result context
* `policies` (List(Dynamic), required) - List of alert policies to be updated
