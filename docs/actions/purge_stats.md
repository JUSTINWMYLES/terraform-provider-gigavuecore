---
page_title: "gigavuecore_purge_stats Action - gigavuecore"
subcategory: ""
description: |-
  Purge Time Series(Deprecated)
---

# gigavuecore_purge_stats Action

Purge Time Series(Deprecated)

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_purge_stats" "example" {
  config {
    start_date = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `start_date` (String, required) - Time Series older than startDate will be purged. In ISO 8601 format.


