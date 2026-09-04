---
page_title: "gigavuecore_image_server Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Image Server
---

# gigavuecore_image_server Resource

Create a new Image Server

## Example Usage

```terraform
resource "gigavuecore_image_server" "example" {
  address  = "example"
  alias    = "example"
  type     = "scp"
  user_pwd = "example"
  username = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `address` (String, required)
* `alias` (String, required) - unique alias for this image server
* `type` (String, required)
* `user_pwd` (String, optional) - user password to use for server login. not applicable for 'tftp'
* `username` (String, optional) - username to use for server login. not applicable for 'tftp'

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
terraform import gigavuecore_image_server.example {alias}
```
