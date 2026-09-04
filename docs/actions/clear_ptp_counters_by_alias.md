---
page_title: "gigavuecore_clear_ptp_counters_by_alias Action - gigavuecore"
subcategory: ""
description: |-
  clear PTP counters by alias
---

# gigavuecore_clear_ptp_counters_by_alias Action

clear PTP counters by alias

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_ptp_counters_by_alias" "example" {
  config {
    alias      = "example"
    box_id     = 0
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the time stamping PTP configuration
* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.
* `cluster_id` (String, required) - Target Cluster ID


