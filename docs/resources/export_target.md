---
page_title: "gigavuecore_export_target Resource - gigavuecore"
subcategory: ""
description: |-
  Create External Export Server
---

# gigavuecore_export_target Resource

Create External Export Server

## Example Usage

```terraform
resource "gigavuecore_export_target" "example" {
  alias       = "example"
  auth_type   = "example"
  broker_type = "example"
  host        = "example"
  password    = "example"
  port        = "example"
  user_name   = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias for external export target server
* `auth_type` (String, required) - Auth Type for external export target server
* `broker_type` (String, required) - Broker Type for external export target server
* `host` (String, required) - Host for external export target server
* `password` (String, optional) - Password for external export target server
* `port` (String, required) - Port for external export target server
* `user_name` (String, optional) - Username for external export target server

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
terraform import gigavuecore_export_target.example {alias}
```
