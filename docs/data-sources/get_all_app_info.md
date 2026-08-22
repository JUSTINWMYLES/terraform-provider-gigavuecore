---
page_title: "gigavuecore_get_all_app_info Data Source - gigavuecore"
subcategory: ""
description: |-
  List all appInfo for unified deployment via environment and connection id
---

# gigavuecore_get_all_app_info Data Source

List all appInfo for unified deployment via environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_all_app_info" "example" {
  env_id = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Set(Object({apps, dp_type, event_dev, major_version, minor_version, name, path, platform_type, serial_num, type, vendor})), computed)

