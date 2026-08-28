---
page_title: "gigavuecore_set_rfs_sync_server Action - gigavuecore"
subcategory: ""
description: |-
  Set RFS server configuration
---

# gigavuecore_set_rfs_sync_server Action

Set RFS server configuration

## Example Usage

```terraform
action "gigavuecore_set_rfs_sync_server" "example" {
  config {
    address    = "example"
    alias      = "example"
    cluster_id = "example"
    hours      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `address` (String, required) - IPv4 address of the RFS server
* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
* `hours` (String, optional) - Period in hours of when to sync, 0 = no automatic sync


