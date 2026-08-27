---
page_title: "gigavuecore_purge_audit_log Action - gigavuecore"
subcategory: ""
description: |-
  Purge Audit Log
---

# gigavuecore_purge_audit_log Action

Purge Audit Log

## Example Usage

```terraform
action "gigavuecore_purge_audit_log" "example" {
  config {
    start_date = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `start_date` (String, optional) - startDate to filter by. In ISO 8601 format. Records older than startDate will be purged.


