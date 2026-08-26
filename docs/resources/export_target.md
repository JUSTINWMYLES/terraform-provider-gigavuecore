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
  alias       = null
  auth_type   = null
  broker_type = null
  host        = null
  password    = null
  port        = null
  user_name   = null
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

* `export_target_alias` (String, computed)
* `password` (String, computed) - Password for external export target server
* `user_name` (String, computed) - Username for external export target server


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_export_target.example {export_target_alias}
```
