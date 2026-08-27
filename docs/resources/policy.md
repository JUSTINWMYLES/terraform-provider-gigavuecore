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
  description    = null
  enabled        = null
  name           = null
  then_do        = []
  when_condition = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `description` (String, optional)
* `enabled` (Boolean, optional)
* `name` (String, required) - AV Policy name
* `then_do` (Attributes List, optional) - AND-joined list of instantiated actions to execute when this policy triggers (see [below for nested schema](#nestedatt--then_do))
* `when_condition` (Attributes List, required) - AND-joined list of instantiated criteria conditions to trigger this policy (see [below for nested schema](#nestedatt--when_condition))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed)
* `enabled` (Boolean, computed)
* `policy_id` (String, computed)
* `then_do` (Attributes List, computed) - AND-joined list of instantiated actions to execute when this policy triggers (see [below for nested schema](#nestedatt--then_do))

<a id="nestedatt--then_do"></a>
### Nested Schema for `then_do`

Required:

* `action` (String) - reference to a pre-defined action template
Optional:

* `params` (Attributes List) - key/value criteria parameters to instantiate the referenced template (see [below for nested schema](#nestedatt--then_do--params))
<a id="nestedatt--then_do--params"></a>
### Nested Schema for `then_do.params`

Optional:

* `key` (String)
* `value` (String)
<a id="nestedatt--when_condition"></a>
### Nested Schema for `when_condition`

Required:

* `condition` (String) - reference to a pre-defined condition template
Optional:

* `params` (Attributes List) - key/value criteria parameters to instantiate the referenced template (see [below for nested schema](#nestedatt--when_condition--params))
<a id="nestedatt--when_condition--params"></a>
### Nested Schema for `when_condition.params`

Optional:

* `key` (String)
* `value` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_policy.example {policy_id}
```
