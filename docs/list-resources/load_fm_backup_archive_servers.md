---
page_title: "gigavuecore_load_fm_backup_archive_servers List Resource - gigavuecore"
subcategory: ""
description: |-
  List FM backup archive servers information
---

# gigavuecore_load_fm_backup_archive_servers List Resource

List FM backup archive servers information

## Example Usage

```terraform
list "gigavuecore_load_fm_backup_archive_servers" "example" {
  provider = gigavuecore
  limit    = 100
}

```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `server_alias` (String, computed) - archive server unique alias


