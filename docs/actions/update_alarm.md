---
page_title: "gigavuecore_update_alarm Action - gigavuecore"
subcategory: ""
description: |-
  Update Alarm
---

# gigavuecore_update_alarm Action

Update Alarm

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


