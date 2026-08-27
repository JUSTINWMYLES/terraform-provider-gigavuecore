---
page_title: "gigavuecore_get_avisi_policies Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Active Visibility Policies
---

# gigavuecore_get_avisi_policies Data Source

Get Active Visibility Policies

## Example Usage

```terraform
data "gigavuecore_get_avisi_policies" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `description` (String)
* `enabled` (Boolean)
* `name` (String) - AV Policy name
* `then_do` (Attributes List) - AND-joined list of instantiated actions to execute when this policy triggers (see [below for nested schema](#nestedatt--items--then_do))
* `when_condition` (Attributes List) - AND-joined list of instantiated criteria conditions to trigger this policy (see [below for nested schema](#nestedatt--items--when_condition))
<a id="nestedatt--items--then_do"></a>
### Nested Schema for `items.then_do`

Read-Only:

* `action` (String) - reference to a pre-defined action template
* `params` (Attributes List) - key/value criteria parameters to instantiate the referenced template (see [below for nested schema](#nestedatt--items--then_do--params))
<a id="nestedatt--items--then_do--params"></a>
### Nested Schema for `items.then_do.params`

Read-Only:

* `key` (String)
* `value` (String)
<a id="nestedatt--items--when_condition"></a>
### Nested Schema for `items.when_condition`

Read-Only:

* `condition` (String) - reference to a pre-defined condition template
* `params` (Attributes List) - key/value criteria parameters to instantiate the referenced template (see [below for nested schema](#nestedatt--items--when_condition--params))
<a id="nestedatt--items--when_condition--params"></a>
### Nested Schema for `items.when_condition.params`

Read-Only:

* `key` (String)
* `value` (String)

