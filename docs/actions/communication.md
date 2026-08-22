---
page_title: "gigavuecore_communication Action - gigavuecore"
subcategory: ""
description: |-
  stop/resume communcation by cluster IDs.
---

# gigavuecore_communication Action

stop/resume communcation by cluster IDs.

## Example Usage

```terraform
action "gigavuecore_communication" "example" {
  config {
    cluster_id = "example"
    disconnect = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster ID
* `disconnect` (Bool, required) - The boolean to stop/resume communication to device.
