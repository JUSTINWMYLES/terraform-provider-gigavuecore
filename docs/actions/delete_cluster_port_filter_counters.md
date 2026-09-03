---
page_title: "gigavuecore_delete_cluster_port_filter_counters Action - gigavuecore"
subcategory: ""
description: |-
  Clear Cluster Port Filter Counters
---

# gigavuecore_delete_cluster_port_filter_counters Action

Clear Cluster Port Filter Counters

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_cluster_port_filter_counters" "example" {
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
* `port_id` (String, required) - port id to filter by.


