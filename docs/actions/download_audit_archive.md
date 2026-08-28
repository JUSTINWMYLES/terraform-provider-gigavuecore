---
page_title: "gigavuecore_download_audit_archive Action - gigavuecore"
subcategory: ""
description: |-
  Archive Audit Log
---

# gigavuecore_download_audit_archive Action

Archive Audit Log

## Example Usage

```terraform
action "gigavuecore_download_audit_archive" "example" {
  config {
    destination = "example"
    purge       = true
    scope       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `destination` (Dynamic, required)
* `purge` (Boolean, optional) - If true, the archived audit log entries will be purged
* `scope` (Dynamic, required)


