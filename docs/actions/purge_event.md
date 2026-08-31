---
page_title: "gigavuecore_purge_event Action - gigavuecore"
subcategory: ""
description: |-
  Purge Event
---

# gigavuecore_purge_event Action

Purge Event

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_purge_event" "example" {
  config {
    start_date = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `start_date` (String, optional) - startDate to filter by. In ISO 8601 format. Records older than startDate will be purged.


