---
page_title: "gigavuecore_load_spine_link_all Data Source - gigavuecore"
subcategory: ""
description: |-
  Load spine-link configuration
---

# gigavuecore_load_spine_link_all Data Source

Load spine-link configuration

## Example Usage

```terraform
data "gigavuecore_load_spine_link_all" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, comment, leaf_box_id, links})), computed)

