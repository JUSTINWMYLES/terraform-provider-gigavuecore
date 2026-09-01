---
page_title: "gigavuecore_redefine_gs_group_gs_group_system_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's System Params
---

# gigavuecore_redefine_gs_group_gs_group_system_params Action

Redefine GS Group's System Params

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gs_group_system_params" "example" {
  config {
    alias                    = "example"
    cluster_id               = "example"
    cpu_load_alarm_threshold = 20
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `cpu_load_alarm_threshold` (Number, optional) - CPU rising threshold percentage 20-99%. Once this threshold is crossed for 5 seconds, an alarm is generated


