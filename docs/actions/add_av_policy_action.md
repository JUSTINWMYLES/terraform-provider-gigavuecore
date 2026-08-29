---
page_title: "gigavuecore_add_av_policy_action Action - gigavuecore"
subcategory: ""
description: |-
  Add Active Visibility Policy Action
---

# gigavuecore_add_av_policy_action Action

Add Active Visibility Policy Action

## Example Usage

```terraform
action "gigavuecore_add_av_policy_action" "example" {
  config {
    action = "example"
    params = [{
      key   = "example"
      value = "example"
    }]
    policy_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `action` (String, required) - reference to a pre-defined action template
* `params` (Attributes List, optional) - key/value criteria parameters to instantiate the referenced template (see [below for nested schema](#nestedatt--params))
* `policy_id` (String, required) - Policy Id

<a id="nestedatt--params"></a>
### Nested Schema for `params`

Optional:

* `key` (String)
* `value` (String)

