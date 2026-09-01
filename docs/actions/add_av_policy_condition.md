---
page_title: "gigavuecore_add_av_policy_condition Action - gigavuecore"
subcategory: ""
description: |-
  Add Active Visibility Policy Condition
---

# gigavuecore_add_av_policy_condition Action

Add Active Visibility Policy Condition

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_add_av_policy_condition" "example" {
  config {
    condition = "example"
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

* `condition` (String, required) - reference to a pre-defined condition template
* `params` (Attributes List, optional) - key/value criteria parameters to instantiate the referenced template (see [below for nested schema](#nestedatt--params))
* `policy_id` (String, required) - Policy Id

<a id="nestedatt--params"></a>
### Nested Schema for `params`

Optional:

* `key` (String)
* `value` (String)

