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

* `yaml_policies` (List(Object({criteria_bindings, name, packet_transformation_bindings, priority, rules, source_bindings, sources, tags})), computed)

