---
page_title: "gigavuecore_get_system_host_banner Data Source - gigavuecore"
subcategory: ""
description: |-
  get system banner configurations
---

# gigavuecore_get_system_host_banner Data Source

get system banner configurations

## Example Usage

```terraform
data "gigavuecore_get_system_host_banner" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Gigamon cluster Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `failed_login_attempts` (String, computed) - failed LoginAttempts
* `login_msg` (String, computed) - Login Message
* `motd` (String, computed) - Message Of The Day

