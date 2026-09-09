---
page_title: "gigavuecore_clear_all_ptp_counters Action - gigavuecore"
subcategory: ""
description: |-
  clear all PTP counters
---

# gigavuecore_clear_all_ptp_counters Action

clear all PTP counters

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


