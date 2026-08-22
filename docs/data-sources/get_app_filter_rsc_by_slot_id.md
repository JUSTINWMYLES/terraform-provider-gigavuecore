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
  cluster_id = null
  slot_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `slot_id` (String, required) - Device slotId

### Attributes

In addition to all arguments above, the following attributes are exported:

* `app_filter_resources` (List(Object({app, free_pool_used, rsvd, rsvd_used})), computed)

