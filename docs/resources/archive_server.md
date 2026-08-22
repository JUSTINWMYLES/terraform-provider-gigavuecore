---
page_title: "gigavuecore_archive_server Resource - gigavuecore"
subcategory: ""
description: |-
  Find FM Backup Archive Server by alias
---

# gigavuecore_archive_server Resource

Find FM Backup Archive Server by alias

## Example Usage

```terraform
resource "gigavuecore_archive_server" "example" {
  address = null
  alias = null
  remote_base_path = null
  type = null
  user_name = null
  user_pwd = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `address` (String, required)
* `alias` (String, required) - archive server unique alias
* `remote_base_path` (String, required) - remote file staging location where FM copies/lists the archived file(s)
* `type` (String, required)
* `user_name` (String, required) - username to use for server login.
* `user_pwd` (String, optional) - user password to use for server login.

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `server_alias` (String, computed)
* `user_pwd` (String, computed) - user password to use for server login.

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_archive_server.example {server_alias}
```
