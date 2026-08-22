---
page_title: "gigavuecore_clear_nrt_stats_counters Action - gigavuecore"
subcategory: ""
description: |-
  Clears Registered flexinline solution from  NRT
---

# gigavuecore_clear_nrt_stats_counters Action

Clears Registered flexinline solution from  NRT

## Example Usage

```terraform
action "gigavuecore_clear_nrt_stats_counters" "example" {
  config {
    solution_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - Clears  flexinline solution from  NRT
