---
page_title: "gigavuecore_update_icap_nrt_stats_config Action - gigavuecore"
subcategory: ""
description: |-
  Register Icap solution to NRT stats
---

# gigavuecore_update_icap_nrt_stats_config Action

Register Icap solution to NRT stats

## Example Usage

```terraform
action "gigavuecore_update_icap_nrt_stats_config" "example" {
  config {
    operation_type = "example"
    solution_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `operation_type` (String, required) - adds icap solution to NRT
* `solution_alias` (String, required) - Register icap solution to NRT
