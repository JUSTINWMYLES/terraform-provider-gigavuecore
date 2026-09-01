---
page_title: "gigavuecore_modify_event_notification_config Action - gigavuecore"
subcategory: ""
description: |-
  Modify Event notification config
---

# gigavuecore_modify_event_notification_config Action

Modify Event notification config

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_modify_event_notification_config" "example" {
  config {
    device_ip = "example"
    enabled   = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `device_ip` (String, required) - Device IP whose event configuration is being changed
* `enabled` (Boolean, required) - Should events be streamed from this device


