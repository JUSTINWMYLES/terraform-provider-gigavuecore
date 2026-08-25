---
page_title: "gigavuecore_delete_alarm Action - gigavuecore"
subcategory: ""
description: |-
  Delete Alarm
---

# gigavuecore_delete_alarm Action

Delete Alarm

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

* `alarm_id` (String, required)


