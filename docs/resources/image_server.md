---
page_title: "gigavuecore_image_server Resource - gigavuecore"
subcategory: ""
description: |-
  Find Image Server by address
---

# gigavuecore_image_server Resource

Find Image Server by address

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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_image_server.example {alias}
```
