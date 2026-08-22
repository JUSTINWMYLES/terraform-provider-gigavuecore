---
page_title: "gigavuecore_load_nrt_stats_flex_inline_solution Data Source - gigavuecore"
subcategory: ""
description: |-
  get NRT stats for Registered flex Inline solution
---

# gigavuecore_load_nrt_stats_flex_inline_solution Data Source

get NRT stats for Registered flex Inline solution

## Example Usage

```terraform
data "gigavuecore_load_nrt_stats_flex_inline_solution" "example" {
  solution_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - get NRT stats for registered flexInline solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, cluster_name, component_type, nrt_alias, stats_data})), computed)

