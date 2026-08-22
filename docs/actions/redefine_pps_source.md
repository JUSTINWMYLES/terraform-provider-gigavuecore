---
page_title: "gigavuecore_redefine_pps_source Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Pulse-Per-Second Source
---

# gigavuecore_redefine_pps_source Action

Redefine Pulse-Per-Second Source

## Example Usage

```terraform
action "gigavuecore_redefine_pps_source" "example" {
  config {
    cluster_id = "example"
    pps_offset = 1
    pps_source_admin = "example"
    pps_source_oper = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `pps_offset` (Number, optional) - in nano-seconds. Configured values are in the range of \[1..280\]. Value of 0 represents the system default
* `pps_source_admin` (String, optional) - Admin-configured timestamp source. Each of the values corresponds to a connector on the HCCv2 Control Card
* `pps_source_oper` (String, optional) - Read-only. The actual runtime timestamp source being used
