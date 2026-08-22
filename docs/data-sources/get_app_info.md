---
page_title: "gigavuecore_get_app_info Data Source - gigavuecore"
subcategory: ""
description: |-
  Get an appInfo for unified deployment via environment and connection id
---

# gigavuecore_get_app_info Data Source

Get an appInfo for unified deployment via environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_app_info" "example" {
  appinfo = null
  env_id = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `appinfo` (String, required) - appinfo identifier
* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `apps` (Object({max_nr_aeps, mem_footprints, name, packet_permission}), computed) - Unified Resource app info for specific app
  * `max_nr_aeps` (String, computed)
  * `mem_footprints` (Object({base, scale}), computed) - app specific memory footprint info
    * `base` (Number, computed)
    * `scale` (Object({item, memory, shared}), computed) - app specific memory footprint scale info
      * `item` (String, computed)
      * `memory` (Number, computed)
      * `shared` (Bool, computed)
  * `name` (String, computed)
  * `packet_permission` (String, computed)
* `dp_type` (String, computed)
* `event_dev` (String, computed)
* `major_version` (Number, computed)
* `minor_version` (Number, computed)
* `name` (String, computed)
* `path` (String, computed)
* `platform_type` (String, computed)
* `serial_num` (Number, computed)
* `type` (String, computed)
* `vendor` (String, computed)

