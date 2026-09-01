---
page_title: "gigavuecore_delete_alarm Action - gigavuecore"
subcategory: ""
description: |-
  Delete Alarm
---

# gigavuecore_delete_alarm Action

Delete Alarm

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_alarm" "example" {
  config {
    alarm_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alarm_id` (String, required) - ID of the target Alarm


