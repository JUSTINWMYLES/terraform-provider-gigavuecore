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
  address          = "example"
  alias            = "example"
  remote_base_path = "example"
  type             = "scp"
  user_name        = "example"
  user_pwd         = "example"
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_archive_server.example {alias}
```
