---
page_title: "gigavuecore_archive_server Resource - gigavuecore"
subcategory: ""
description: |-
  Create FM backup archive server spec
---

# gigavuecore_archive_server Resource

Create FM backup archive server spec

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_archive_server.example {alias}
```
