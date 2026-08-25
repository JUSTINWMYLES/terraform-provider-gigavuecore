---
page_title: "gigavuecore_download_syslog_archive Action - gigavuecore"
subcategory: ""
description: |-
  Archive Syslogs
---

# gigavuecore_download_syslog_archive Action

Archive Syslogs

## Example Usage

```terraform
action "gigavuecore_download_syslog_archive" "example" {
  config {
    cluster_id  = "example"
    destination = "example"
    purge       = true
    scope       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - ID of the cluster whose logs are being archived
* `destination` (Dynamic, required)
* `purge` (Boolean, optional) - If true, the archived event entries will be purged
* `scope` (Dynamic, required)


