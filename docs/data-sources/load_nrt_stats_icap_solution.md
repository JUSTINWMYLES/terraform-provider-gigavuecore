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
  solution_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - get NRT stats for registered icap solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `cluster_name` (String)
* `component_type` (String)
* `stats_data` (Attributes) (see [below for nested schema](#nestedatt--items--stats_data))
<a id="nestedatt--items--stats_data"></a>
### Nested Schema for `items.stats_data`

Read-Only:

* `alias` (String)
* `real_time_stats` (Attributes) (see [below for nested schema](#nestedatt--items--stats_data--real_time_stats))
<a id="nestedatt--items--stats_data--real_time_stats"></a>
### Nested Schema for `items.stats_data.real_time_stats`

Read-Only:

* `rate_ato_b` (Attributes) (see [below for nested schema](#nestedatt--items--stats_data--real_time_stats--rate_ato_b))
* `rate_bto_a` (Attributes) (see [below for nested schema](#nestedatt--items--stats_data--real_time_stats--rate_bto_a))
<a id="nestedatt--items--stats_data--real_time_stats--rate_ato_b"></a>
### Nested Schema for `items.stats_data.real_time_stats.rate_ato_b`

Read-Only:

* `rx` (Number)
* `tx` (Number)
<a id="nestedatt--items--stats_data--real_time_stats--rate_bto_a"></a>
### Nested Schema for `items.stats_data.real_time_stats.rate_bto_a`

Read-Only:

* `rx` (Number)
* `tx` (Number)

