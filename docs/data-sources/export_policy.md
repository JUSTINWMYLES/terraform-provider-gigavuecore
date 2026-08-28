---
page_title: "gigavuecore_export_policy Data Source - gigavuecore"
subcategory: ""
description: |-
  Export policies
---

# gigavuecore_export_policy Data Source

Export policies

## Example Usage

```terraform
data "gigavuecore_export_policy" "example" {
  name = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `name` (String, optional) - policy name

### Attributes

In addition to all arguments above, the following attributes are exported:

* `yaml_policies` (Attributes List, computed) (see [below for nested schema](#nestedatt--yaml_policies))

<a id="nestedatt--yaml_policies"></a>
### Nested Schema for `yaml_policies`

Read-Only:

* `criteria_bindings` (Map of List of Object)
* `name` (String)
* `packet_transformation_bindings` (Map of List of Object)
* `priority` (Boolean)
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--rules))
* `source_bindings` (Map of List of Object)
* `sources` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--sources))
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--tags))
<a id="nestedatt--yaml_policies--rules"></a>
### Nested Schema for `yaml_policies.rules`

Read-Only:

* `criteria` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--rules--criteria))
* `high_priority_drop` (Boolean)
* `no_expansion_tags` (List of String)
* `packet_transformation` (Attributes) (see [below for nested schema](#nestedatt--yaml_policies--rules--packet_transformation))
* `rule_name` (String)
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--rules--tags))
* `tools` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--rules--tools))
* `type` (String)
<a id="nestedatt--yaml_policies--rules--criteria"></a>
### Nested Schema for `yaml_policies.rules.criteria`

Read-Only:

* `criteria_name` (String)
* `filters` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--rules--criteria--filters))
<a id="nestedatt--yaml_policies--rules--criteria--filters"></a>
### Nested Schema for `yaml_policies.rules.criteria.filters`

Read-Only:

* `properties` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--rules--criteria--filters--properties))
* `type` (String)
* `values` (List of String)
<a id="nestedatt--yaml_policies--rules--criteria--filters--properties"></a>
### Nested Schema for `yaml_policies.rules.criteria.filters.properties`

Read-Only:

* `key` (String)
* `values` (List of Dynamic)
<a id="nestedatt--yaml_policies--rules--packet_transformation"></a>
### Nested Schema for `yaml_policies.rules.packet_transformation`

Read-Only:

* `apps` (Attributes List) (see [below for nested schema](#nestedatt--yaml_policies--rules--packet_transformation--apps))
* `engine_ports` (List of String)
<a id="nestedatt--yaml_policies--rules--packet_transformation--apps"></a>
### Nested Schema for `yaml_policies.rules.packet_transformation.apps`

Read-Only:

* `parameters` (Dynamic)
* `type` (String)
<a id="nestedatt--yaml_policies--rules--tags"></a>
### Nested Schema for `yaml_policies.rules.tags`

Read-Only:

* `key` (String)
* `values` (List of String)
<a id="nestedatt--yaml_policies--rules--tools"></a>
### Nested Schema for `yaml_policies.rules.tools`

Read-Only:

* `ports` (List of String)
* `type` (String)
<a id="nestedatt--yaml_policies--sources"></a>
### Nested Schema for `yaml_policies.sources`

Read-Only:

* `ports` (List of String)
* `type` (String)
<a id="nestedatt--yaml_policies--tags"></a>
### Nested Schema for `yaml_policies.tags`

Read-Only:

* `key` (String)
* `values` (List of String)

