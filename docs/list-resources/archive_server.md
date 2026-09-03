---
page_title: "gigavuecore_archive_server List Resource - gigavuecore"
subcategory: ""
description: |-
  List FM backup archive servers information
---

# gigavuecore_archive_server List Resource

List FM backup archive servers information

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_archive_server" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `server_alias` (String, computed) - archive server unique alias


