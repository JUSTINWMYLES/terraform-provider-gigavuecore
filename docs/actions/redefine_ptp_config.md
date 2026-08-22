---
page_title: "gigavuecore_redefine_ptp_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine PTP Configuration
---

# gigavuecore_redefine_ptp_config Action

Redefine PTP Configuration

## Example Usage

```terraform
action "gigavuecore_redefine_ptp_config" "example" {
  config {
    cluster_id = "example"
    enabled = true
    mode = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Bool, required) - enable/disable use of NTP for synchronization of the system's clock
* `mode` (String, optional)
