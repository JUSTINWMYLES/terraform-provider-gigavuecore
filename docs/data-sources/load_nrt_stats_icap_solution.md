---
page_title: "gigavuecore_load_nrt_stats_icap_solution Data Source - gigavuecore"
subcategory: ""
description: |-
  get NRT stats for Registered Icap solution
---

# gigavuecore_load_nrt_stats_icap_solution Data Source

get NRT stats for Registered Icap solution

## Example Usage

```terraform
data "gigavuecore_load_nrt_stats_icap_solution" "example" {
  solution_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - get NRT stats for registered icap solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, cluster_name, component_type, stats_data})), computed)

