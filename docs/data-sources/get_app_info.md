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
  appinfo  = "example"
  env_id   = "example"
  unify_id = "example"
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

* `apps` (Attributes, computed) - Unified Resource app info for specific app (see [below for nested schema](#nestedatt--apps))
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

<a id="nestedatt--apps"></a>
### Nested Schema for `apps`

Read-Only:

* `max_nr_aeps` (String)
* `mem_footprints` (Attributes) - app specific memory footprint info (see [below for nested schema](#nestedatt--apps--mem_footprints))
* `name` (String)
* `packet_permission` (String)
<a id="nestedatt--apps--mem_footprints"></a>
### Nested Schema for `apps.mem_footprints`

Read-Only:

* `base` (Number)
* `scale` (Attributes) - app specific memory footprint scale info (see [below for nested schema](#nestedatt--apps--mem_footprints--scale))
<a id="nestedatt--apps--mem_footprints--scale"></a>
### Nested Schema for `apps.mem_footprints.scale`

Read-Only:

* `item` (String)
* `memory` (Number)
* `shared` (Boolean)

