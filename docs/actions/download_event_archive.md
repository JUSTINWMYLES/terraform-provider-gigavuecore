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
    destination = "example"
    purge = true
    scope = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `destination` (Dynamic, required)
* `purge` (Bool, optional) - If true, the archived event entries will be purged
* `scope` (Dynamic, required)
