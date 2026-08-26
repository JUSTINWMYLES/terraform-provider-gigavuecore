---
page_title: "gigavuecore_load_pps_source Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Pulse-Per-Second Source config
---

# gigavuecore_load_pps_source Data Source

Load Pulse-Per-Second Source config

## Example Usage

```terraform
data "gigavuecore_load_pps_source" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `pps_offset` (Number, computed) - in nano-seconds. Configured values are in the range of \[1..280\]. Value of 0 represents the system default
* `pps_source_admin` (String, computed) - Admin-configured timestamp source. Each of the values corresponds to a connector on the HCCv2 Control Card
* `pps_source_oper` (String, computed) - Read-only. The actual runtime timestamp source being used


