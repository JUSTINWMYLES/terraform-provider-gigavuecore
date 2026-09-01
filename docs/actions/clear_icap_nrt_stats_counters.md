---
page_title: "gigavuecore_clear_icap_nrt_stats_counters Action - gigavuecore"
subcategory: ""
description: |-
  Clears Registered icap solution Stats
---

# gigavuecore_clear_icap_nrt_stats_counters Action

Clears Registered icap solution Stats

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_icap_nrt_stats_counters" "example" {
  config {
    solution_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - Clears icap solution Stats


