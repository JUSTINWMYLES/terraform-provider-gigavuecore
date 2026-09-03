---
page_title: "gigavuecore_acknowledge_multiple_alarms Action - gigavuecore"
subcategory: ""
description: |-
  Acknowledge Multiple Alarms
---

# gigavuecore_acknowledge_multiple_alarms Action

Acknowledge Multiple Alarms

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_acknowledge_multiple_alarms" "example" {
  config {
    alarm_ids = ["example"]
    comment   = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alarm_ids` (List of String, required) - IDs of alarms that needs to be acknowledged
* `comment` (String, optional) - Alarm comment


