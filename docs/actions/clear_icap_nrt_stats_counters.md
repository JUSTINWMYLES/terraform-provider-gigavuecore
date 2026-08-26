---
page_title: "gigavuecore_clear_icap_nrt_stats_counters Action - gigavuecore"
subcategory: ""
description: |-
  Clears Registered icap solution Stats
---

# gigavuecore_clear_icap_nrt_stats_counters Action

Clears Registered icap solution Stats

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


