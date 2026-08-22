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
  alias = null
  cluster_id = null
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
* `rfs_sync_info` (List(Object({key_name, key_token})), computed)
* `sync_period` (String, computed) - Period in hours of when to sync, 0 = no automatic sync

