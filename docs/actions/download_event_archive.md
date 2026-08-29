---
page_title: "gigavuecore_download_event_archive Action - gigavuecore"
subcategory: ""
description: |-
  Archive Events
---

# gigavuecore_download_event_archive Action

Archive Events

## Example Usage

```terraform
action "gigavuecore_download_event_archive" "example" {
  config {
    destination = {
      sftp = {
        file_path    = "example"
        host_address = "example"
        password     = "example"
        username     = "example"
      }
    }
    purge = true
    scope = {
      start_date = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `destination` (Attributes, required) (see [below for nested schema](#nestedatt--destination))
* `purge` (Boolean, optional) - If true, the archived event entries will be purged
* `scope` (Attributes, required) (see [below for nested schema](#nestedatt--scope))

<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Required:

* `sftp` (Attributes) (see [below for nested schema](#nestedatt--destination--sftp))
<a id="nestedatt--destination--sftp"></a>
### Nested Schema for `destination.sftp`

Required:

* `file_path` (String) - specifies the path on server to store the archive. Example: '/root/dir/archive.zip'
* `host_address` (String) - Specifies the host address of server to save the archive
* `password` (String) - specifies the password to use for server login
* `username` (String) - specifies the username for the server login
<a id="nestedatt--scope"></a>
### Nested Schema for `scope`

Optional:

* `start_date` (String) - Records older than startDate will be archived. In ISO 8601 format.

