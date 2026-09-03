---
page_title: "gigavuecore_update_alarm Action - gigavuecore"
subcategory: ""
description: |-
  Update Alarm
---

# gigavuecore_update_alarm Action

Update Alarm

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_alarm" "example" {
  config {
    acknowledged = true
    alarm_id     = "example"
    comment      = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `acknowledged` (Boolean, optional) - Acknowledgement status
* `alarm_id` (String, required) - ID of the target Alarm
* `comment` (String, optional) - Alarm comments


