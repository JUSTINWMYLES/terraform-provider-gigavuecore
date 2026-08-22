---
page_title: "gigavuecore_backup_fm_config Action - gigavuecore"
subcategory: ""
description: |-
  Backup FM configuration state
---

# gigavuecore_backup_fm_config Action

Backup FM configuration state

## Example Usage

```terraform
action "gigavuecore_backup_fm_config" "example" {
  config {
    archive_server = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `archive_server` (String, required) - Alias of the target Archive Server
