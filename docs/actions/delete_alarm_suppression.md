---
page_title: "gigavuecore_delete_alarm_suppression Action - gigavuecore"
subcategory: ""
description: |-
  Delete multiple alarm suppression rules
---

# gigavuecore_delete_alarm_suppression Action

Delete multiple alarm suppression rules

## Example Usage

```terraform
action "gigavuecore_delete_alarm_suppression" "example" {
  config {
    suppressed_entities = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `suppressed_entities` (List(Dynamic), required) - Suppression Rules for resource
