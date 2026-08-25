---
page_title: "gigavuecore_purge_stats Action - gigavuecore"
subcategory: ""
description: |-
  Purge Time Series(Deprecated)
---

# gigavuecore_purge_stats Action

Purge Time Series(Deprecated)

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


