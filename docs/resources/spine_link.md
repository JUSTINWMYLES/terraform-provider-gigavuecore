---
page_title: "gigavuecore_spine_link Resource - gigavuecore"
subcategory: ""
description: |-
  Load spine-link
---

# gigavuecore_spine_link Resource

Load spine-link

## Example Usage

```terraform
resource "gigavuecore_spine_link" "example" {
  alias = null
  comment = null
  gigastreams = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `comment` (String, optional)
* `gigastreams` (List(String), required) - leaf node stack gigastream aliases

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `leaf_box_id` (Number, computed) - box-id of the leaf node this spineLink is configured on
* `links` (List(Object({leaf_gigastream, spine_box_id, spine_gigastream, stack_link})), computed) - List of links from this Leaf node to the Spine Nodes

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_spine_link.example {alias}
```
