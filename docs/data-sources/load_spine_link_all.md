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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `comment` (String)
* `leaf_box_id` (Number) - box-id of the leaf node this spineLink is configured on
* `links` (Attributes List) - List of links from this Leaf node to the Spine Nodes (see [below for nested schema](#nestedatt--items--links))
<a id="nestedatt--items--links"></a>
### Nested Schema for `items.links`

Read-Only:

* `leaf_gigastream` (String) - leaf node stack gigastream alias
* `spine_box_id` (Number) - box-id of the spine nodes this spineLink is connected to
* `spine_gigastream` (String) - spine node gigastream alias
* `stack_link` (String) - Alias of the Stack Link this Spine Link is running over. Will be empty if corresponding StackLink is not yet created

