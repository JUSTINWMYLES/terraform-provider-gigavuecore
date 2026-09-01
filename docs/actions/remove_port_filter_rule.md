---
page_title: "gigavuecore_remove_port_filter_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a filtering rule from a Port
---

# gigavuecore_remove_port_filter_rule Action

Remove a filtering rule from a Port

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_remove_port_filter_rule" "example" {
  config {
    cluster_id = "example"
    port_id    = "example"
    rule_id    = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port_id` (String, required) - id of the target device Port (format: boxId\_slotId\_port, example: 1\_1\_c1)
* `rule_id` (Number, required) - filter rule id


