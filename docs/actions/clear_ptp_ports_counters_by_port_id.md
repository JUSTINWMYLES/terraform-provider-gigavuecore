---
page_title: "gigavuecore_clear_ptp_ports_counters_by_port_id Action - gigavuecore"
subcategory: ""
description: |-
  clear PTP port counters by port ID
---

# gigavuecore_clear_ptp_ports_counters_by_port_id Action

clear PTP port counters by port ID

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_ptp_ports_counters_by_port_id" "example" {
  config {
    cluster_id = "example"
    port_id    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_id` (String, required) - id of the target device Port (format: boxId\_slotId\_port, example: 1\_1\_c1)


