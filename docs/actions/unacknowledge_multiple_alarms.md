---
page_title: "gigavuecore_unacknowledge_multiple_alarms Action - gigavuecore"
subcategory: ""
description: |-
  Unacknowledge Multiple Alarms
---

# gigavuecore_unacknowledge_multiple_alarms Action

Unacknowledge Multiple Alarms

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_unacknowledge_multiple_alarms" "example" {
  config {
    alarm_ids = ["example"]
    comment   = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alarm_ids` (List of String, required) - IDs of alarms that needs to be unacknowledged
* `comment` (String, optional) - Alarm comment


