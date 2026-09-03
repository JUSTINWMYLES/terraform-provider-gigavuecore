---
page_title: "gigavuecore_reboot_cluster_nodes Action - gigavuecore"
subcategory: ""
description: |-
  reboot physical cluster nodes
---

# gigavuecore_reboot_cluster_nodes Action

reboot physical cluster nodes

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_reboot_cluster_nodes" "example" {
  config {
    cluster_ids                = ["example"]
    node_ids                   = ["example"]
    skip_not_reachable_devices = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_ids` (List of String, optional) - ids of clusters to reboot. Every node in these clusters will be rebooted
* `node_ids` (List of String, optional) - ids of individual nodes to reboot
* `skip_not_reachable_devices` (Boolean, optional) - indicates whether not-reachable nodes(if any) should be skipped and continue with reboot operation


