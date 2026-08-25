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
  env_id   = null
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

* `items` (Attributes Set, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `apps` (Attributes) - Unified Resource app info for specific app (see [below for nested schema](#nestedatt--items--apps))
* `dp_type` (String)
* `event_dev` (String)
* `major_version` (Number)
* `minor_version` (Number)
* `name` (String)
* `path` (String)
* `platform_type` (String)
* `serial_num` (Number)
* `type` (String)
* `vendor` (String)
<a id="nestedatt--items--apps"></a>
### Nested Schema for `items.apps`

Read-Only:

* `max_nr_aeps` (String)
* `mem_footprints` (Attributes) - app specific memory footprint info (see [below for nested schema](#nestedatt--items--apps--mem_footprints))
* `name` (String)
* `packet_permission` (String)
<a id="nestedatt--items--apps--mem_footprints"></a>
### Nested Schema for `items.apps.mem_footprints`

Read-Only:

* `base` (Number)
* `scale` (Attributes) - app specific memory footprint scale info (see [below for nested schema](#nestedatt--items--apps--mem_footprints--scale))
<a id="nestedatt--items--apps--mem_footprints--scale"></a>
### Nested Schema for `items.apps.mem_footprints.scale`

Read-Only:

* `item` (String)
* `memory` (Number)
* `shared` (Boolean)

