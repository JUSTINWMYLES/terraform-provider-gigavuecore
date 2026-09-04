---
page_title: "gigavuecore_delete_netflow_exporter_filter Action - gigavuecore"
subcategory: ""
description: |-
  Delete Netflow Exporter Filter
---

# gigavuecore_delete_netflow_exporter_filter Action

Delete Netflow Exporter Filter

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_netflow_exporter_filter" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID


