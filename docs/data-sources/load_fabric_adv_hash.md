---
page_title: "gigavuecore_load_fabric_adv_hash Data Source - gigavuecore"
subcategory: ""
description: |-
  Load fabric advanced hash
---

# gigavuecore_load_fabric_adv_hash Data Source

Load fabric advanced hash

## Example Usage

```terraform
data "gigavuecore_load_fabric_adv_hash" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - device box id. valid range 1 - 64

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (Number) - box id
* `fields` (Set of String) - Required when 'type' == 'fields'. 'gtpteid' requires fields (portsrc and portdst) or (port6src and port6dst)
* `type` (String)

