---
page_title: "gigavuecore_restore_fm_config Action - gigavuecore"
subcategory: ""
description: |-
  Restore backed up FM state
---

# gigavuecore_restore_fm_config Action

Restore backed up FM state

## Example Usage

```terraform
action "gigavuecore_restore_fm_config" "example" {
  config {
    archive_server = "example"
    file_path = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `archive_server` (String, required) - Alias of the target Archive Server
* `file_path` (String, required) - Path of the file
