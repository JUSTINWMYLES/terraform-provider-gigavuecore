---
page_title: "gigavuecore_get_current_physical_apps Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the currently licensed featre/application set for a chassis or card
---

# gigavuecore_get_current_physical_apps Data Source

Gives the currently licensed featre/application set for a chassis or card

## Example Usage

```terraform
data "gigavuecore_get_current_physical_apps" "example" {
  box_id = null
  cluster_name = null
  slot = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - box ID of the device in the cluster
* `cluster_name` (String, required) - name of chassis cluster
* `slot` (String, optional) - slot number in the box (if not provided, chassis-level features returned)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(String), computed)

