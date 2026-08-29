---
page_title: "gigavuecore_bulk_policy_enable Action - gigavuecore"
subcategory: ""
description: |-
  Enable all policies or by policy Ids
---

# gigavuecore_bulk_policy_enable Action

Enable all policies or by policy Ids

## Example Usage

```terraform
action "gigavuecore_bulk_policy_enable" "example" {
  config {
    enabled    = true
    policy_ids = [ "example" ]
    type       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `enabled` (Boolean, required) - enable/disable policies
* `policy_ids` (List of String, optional) - list of policy Ids
* `type` (String, required) - Enable all policies or by policy Ids


