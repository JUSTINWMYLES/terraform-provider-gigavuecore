---
page_title: "gigavuecore_load_fm_licensing_summary Data Source - gigavuecore"
subcategory: ""
description: |-
  Load FM Licensing Summary
---

# gigavuecore_load_fm_licensing_summary Data Source

Load FM Licensing Summary

## Example Usage

```terraform
data "gigavuecore_load_fm_licensing_summary" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `aws` (Attributes, computed) - AWS License summary (see [below for nested schema](#nestedatt--aws))
* `azure` (Attributes, computed) - Azure License summary (see [below for nested schema](#nestedatt--azure))
* `base_bundle` (Attributes, computed) - Gigamon FM License Base Bundle summary (see [below for nested schema](#nestedatt--base_bundle))
* `features` (Attributes, computed) - Gigamon FM License Features (see [below for nested schema](#nestedatt--features))
* `open_stack` (Attributes, computed) - OpenStack License summary (see [below for nested schema](#nestedatt--open_stack))
* `target_fm_id` (String, computed)
* `vmm` (Attributes, computed) - VMM License summary (see [below for nested schema](#nestedatt--vmm))

<a id="nestedatt--aws"></a>
### Nested Schema for `aws`

Read-Only:

* `licensed` (Boolean)
* `v_nics` (Number) - number of licensed vNics

<a id="nestedatt--azure"></a>
### Nested Schema for `azure`

Read-Only:

* `licensed` (Boolean)
* `v_nics` (Number) - number of licensed vNics

<a id="nestedatt--base_bundle"></a>
### Nested Schema for `base_bundle`

Read-Only:

* `nodes` (Number) - number of licensed physical nodes
* `type` (String)

<a id="nestedatt--features"></a>
### Nested Schema for `features`

Read-Only:

* `dashboard_customization` (Attributes) - Gigamon FM License Dashboard Customization Feature summary (see [below for nested schema](#nestedatt--features--dashboard_customization))
* `reporting` (Attributes) - Gigamon FM License Reporting Feature summary (see [below for nested schema](#nestedatt--features--reporting))
* `traffic_analyzer` (Attributes) - Gigamon FM License Traffic Analyzer Feature summary (see [below for nested schema](#nestedatt--features--traffic_analyzer))
* `trending` (Attributes) - Gigamon FM License Trending Feature summary (see [below for nested schema](#nestedatt--features--trending))

<a id="nestedatt--features--dashboard_customization"></a>
### Nested Schema for `features.dashboard_customization`

Read-Only:

* `type` (String)

<a id="nestedatt--features--reporting"></a>
### Nested Schema for `features.reporting`

Read-Only:

* `type` (String)

<a id="nestedatt--features--traffic_analyzer"></a>
### Nested Schema for `features.traffic_analyzer`

Read-Only:

* `type` (String)

<a id="nestedatt--features--trending"></a>
### Nested Schema for `features.trending`

Read-Only:

* `historical_days` (Number)
* `type` (String)

<a id="nestedatt--open_stack"></a>
### Nested Schema for `open_stack`

Read-Only:

* `licensed` (Boolean)
* `v_nics` (Number) - number of licensed vNics

<a id="nestedatt--vmm"></a>
### Nested Schema for `vmm`

Read-Only:

* `historical_days` (Number)
* `licensed` (Boolean)
* `nodes` (Number) - number of licensed GVM nodes

