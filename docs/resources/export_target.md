---
page_title: "gigavuecore_export_target Resource - gigavuecore"
subcategory: ""
description: |-
  Load External Export Server by alias
---

# gigavuecore_export_target Resource

Load External Export Server by alias

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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `password` (String, computed) - Password for external export target server
* `user_name` (String, computed) - Username for external export target server


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_export_target.example {alias}
```
