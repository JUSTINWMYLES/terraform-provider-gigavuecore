---
page_title: "gigavuecore_enable_alarm_suppression Action - gigavuecore"
subcategory: ""
description: |-
  Enable alarm suppression rules for multiple resources
---

# gigavuecore_enable_alarm_suppression Action

Enable alarm suppression rules for multiple resources

## Example Usage

```terraform
action "gigavuecore_enable_alarm_suppression" "example" {
  config {
    suppressed_entities = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `suppressed_entities` (List of Dynamic, required) - Suppression Rules for resource


