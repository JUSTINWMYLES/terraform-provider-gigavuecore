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
  alias       = "example"
  comment     = "example"
  gigastreams = [ "example" ]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `comment` (String, optional)
* `gigastreams` (List of String, required) - leaf node stack gigastream aliases

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `leaf_box_id` (Number, computed) - box-id of the leaf node this spineLink is configured on
* `links` (Attributes List, computed) - List of links from this Leaf node to the Spine Nodes (see [below for nested schema](#nestedatt--links))

<a id="nestedatt--links"></a>
### Nested Schema for `links`

Read-Only:

* `leaf_gigastream` (String) - leaf node stack gigastream alias
* `spine_box_id` (Number) - box-id of the spine nodes this spineLink is connected to
* `spine_gigastream` (String) - spine node gigastream alias
* `stack_link` (String) - Alias of the Stack Link this Spine Link is running over. Will be empty if corresponding StackLink is not yet created

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_spine_link.example {alias}
```
