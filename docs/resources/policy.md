---
page_title: "gigavuecore_policy Resource - gigavuecore"
subcategory: ""
description: |-
  Get Active Visibility Policy
---

# gigavuecore_policy Resource

Get Active Visibility Policy

## Example Usage

```terraform
resource "gigavuecore_policy" "example" {
  description = null
  enabled = null
  name = null
  then_do = []
  when_condition = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional)
* `enabled` (Bool, optional)
* `name` (String, required) - AV Policy name
* `then_do` (List(Object({action, params})), optional) - AND-joined list of instantiated actions to execute when this policy triggers
* `when_condition` (List(Object({condition, params})), required) - AND-joined list of instantiated criteria conditions to trigger this policy

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed)
* `enabled` (Bool, computed)
* `policy_id` (String, computed)
* `then_do` (List(Object({action, params})), computed) - AND-joined list of instantiated actions to execute when this policy triggers

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_policy.example {policy_id}
```
