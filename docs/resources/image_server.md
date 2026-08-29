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
  type     = "example"
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `user_pwd` (String, computed) - user password to use for server login. not applicable for 'tftp'
* `username` (String, computed) - username to use for server login. not applicable for 'tftp'


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_image_server.example {alias}
```
