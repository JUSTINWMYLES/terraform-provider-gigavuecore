---
page_title: "gigavuecore_delete_multiple_alarms Action - gigavuecore"
subcategory: ""
description: |-
  Delete Multiple Alarms
---

# gigavuecore_delete_multiple_alarms Action

Delete Multiple Alarms

## Example Usage

```terraform
action "gigavuecore_delete_multiple_alarms" "example" {
  config {
    alarm_ids = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alarm_ids` (List(String), required) - IDs of alarms that needs to be deleted
