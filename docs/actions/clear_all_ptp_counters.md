---
page_title: "gigavuecore_clear_all_ptp_counters Action - gigavuecore"
subcategory: ""
description: |-
  clear all PTP counters
---

# gigavuecore_clear_all_ptp_counters Action

clear all PTP counters

## Example Usage

```terraform
action "gigavuecore_clear_all_ptp_counters" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
