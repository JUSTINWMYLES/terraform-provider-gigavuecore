---
page_title: "gigavuecore_get_app_filter_rsc_by_slot_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Application Filter Resource by slotId
---

# gigavuecore_get_app_filter_rsc_by_slot_id Data Source

Load Application Filter Resource by slotId

## Example Usage

```terraform
data "gigavuecore_get_app_filter_rsc_by_slot_id" "example" {
  cluster_id = "example"
  slot_id    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `slot_id` (String, required) - Device slotId

### Attributes

In addition to all arguments above, the following attributes are exported:

* `app_filter_resources` (Attributes List, computed) (see [below for nested schema](#nestedatt--app_filter_resources))

<a id="nestedatt--app_filter_resources"></a>
### Nested Schema for `app_filter_resources`

Read-Only:

* `app` (String) - vfp client application
* `free_pool_used` (Number) - vfp resource used from free pool
* `rsvd` (Number) - vfp resource reserved for app
* `rsvd_used` (Number) - resource used from 'rsvd'

