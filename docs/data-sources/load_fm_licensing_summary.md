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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `aws` (Object({licensed, v_nics}), computed) - AWS License summary
  * `licensed` (Bool, computed)
  * `v_nics` (Number, computed) - number of licensed vNics
* `azure` (Object({licensed, v_nics}), computed) - Azure License summary
  * `licensed` (Bool, computed)
  * `v_nics` (Number, computed) - number of licensed vNics
* `base_bundle` (Object({nodes, type}), computed) - Gigamon FM License Base Bundle summary
  * `nodes` (Number, computed) - number of licensed physical nodes
  * `type` (String, computed)
* `features` (Object({dashboard_customization, reporting, traffic_analyzer, trending}), computed) - Gigamon FM License Features
  * `dashboard_customization` (Object({type}), computed) - Gigamon FM License Dashboard Customization Feature summary
    * `type` (String, computed)
  * `reporting` (Object({type}), computed) - Gigamon FM License Reporting Feature summary
    * `type` (String, computed)
  * `traffic_analyzer` (Object({type}), computed) - Gigamon FM License Traffic Analyzer Feature summary
    * `type` (String, computed)
  * `trending` (Object({historical_days, type}), computed) - Gigamon FM License Trending Feature summary
    * `historical_days` (Number, computed)
    * `type` (String, computed)
* `open_stack` (Object({licensed, v_nics}), computed) - OpenStack License summary
  * `licensed` (Bool, computed)
  * `v_nics` (Number, computed) - number of licensed vNics
* `target_fm_id` (String, computed)
* `vmm` (Object({historical_days, licensed, nodes}), computed) - VMM License summary
  * `historical_days` (Number, computed)
  * `licensed` (Bool, computed)
  * `nodes` (Number, computed) - number of licensed GVM nodes

