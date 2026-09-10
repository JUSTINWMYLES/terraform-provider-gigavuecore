---
page_title: "gigavuecore_policy Resource - gigavuecore"
subcategory: ""
description: |-
  Create Active Visibility Policy
---

# gigavuecore_policy Resource

Create Active Visibility Policy

## Example Usage

```terraform
resource "gigavuecore_policy" "example" {
  description = "example"
  enabled     = true
  name        = "example"
  then_do = [{
    action = "example"
    params = [{
      key   = "example"
      value = "example"
    }]
  }]
  when_condition = [{
    condition = "example"
    params = [{
      key   = "example"
      value = "example"
    }]
  }]
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

* `policy_id` (String, computed) - Policy Id

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_policy.example {policy_id}
```
