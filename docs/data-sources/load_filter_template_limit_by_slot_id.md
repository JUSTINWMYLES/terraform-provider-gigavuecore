---
page_title: "gigavuecore_load_filter_template_limit_by_slot_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Filter Template Limits by slot Id
---

# gigavuecore_load_filter_template_limit_by_slot_id Data Source

Load Filter Template Limits by slot Id

## Example Usage

```terraform
data "gigavuecore_load_filter_template_limit_by_slot_id" "example" {
  cluster_id = null
  slot_id    = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `slot_id` (String, required) - Device card slot Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `filter_limits` (Attributes List, computed) (see [below for nested schema](#nestedatt--filter_limits))

<a id="nestedatt--filter_limits"></a>
### Nested Schema for `filter_limits`

Read-Only:

* `alias` (String) - alias of filter template
* `limit` (Number)
* `qualifiers` (List of String)

