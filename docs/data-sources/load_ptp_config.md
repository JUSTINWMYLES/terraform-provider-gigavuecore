---
page_title: "gigavuecore_load_ptp_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load PTP Configuration
---

# gigavuecore_load_ptp_config Data Source

Load PTP Configuration

## Example Usage

```terraform
data "gigavuecore_load_ptp_config" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `enabled` (Boolean, computed) - enable/disable use of NTP for synchronization of the system's clock
* `mode` (String, computed)


