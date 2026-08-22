---
page_title: "gigavuecore_edit_alarm_suppression Action - gigavuecore"
subcategory: ""
description: |-
  Edit multiple alarm suppression rules
---

# gigavuecore_edit_alarm_suppression Action

Edit multiple alarm suppression rules

## Example Usage

```terraform
action "gigavuecore_edit_alarm_suppression" "example" {
  config {
    suppressed_entities = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `suppressed_entities` (List(Dynamic), required) - Suppression Rules for resource
