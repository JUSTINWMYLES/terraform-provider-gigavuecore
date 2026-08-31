---
page_title: "gigavuecore_get_rfs_sync Data Source - gigavuecore"
subcategory: ""
description: |-
  Get HSM Rfs Sync
---

# gigavuecore_get_rfs_sync Data Source

Get HSM Rfs Sync

## Example Usage

```terraform
data "gigavuecore_get_rfs_sync" "example" {
  alias      = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - id of the defining cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `last_sync` (String, computed) - Time that the last sync happened
* `last_sync_method` (String, computed) - How the the last sync happened,either auto or manuel
* `next_sync` (String, computed) - Time that next sync will happen
* `rfs_address` (String, computed) - IPv4 address of the RFS server
* `rfs_sync_info` (Attributes List, computed) (see [below for nested schema](#nestedatt--rfs_sync_info))
* `sync_period` (String, computed) - Period in hours of when to sync, 0 = no automatic sync

<a id="nestedatt--rfs_sync_info"></a>
### Nested Schema for `rfs_sync_info`

Read-Only:

* `key_name` (String) - Key Name
* `key_token` (String) - Key Token

