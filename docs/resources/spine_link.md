---
page_title: "gigavuecore_spine_link Resource - gigavuecore"
subcategory: ""
description: |-
  Add spine-link configuration
---

# gigavuecore_spine_link Resource

Add spine-link configuration

## Example Usage

```terraform
resource "gigavuecore_spine_link" "example" {
  alias       = "example"
  comment     = "example"
  gigastreams = ["example"]
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

* `leaf_box_id` (Number, computed) - box-id of the leaf node this spineLink is configured on
* `links` (Attributes List, computed) - List of links from this Leaf node to the Spine Nodes (see [below for nested schema](#nestedatt--links))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--links"></a>
### Nested Schema for `links`

Read-Only:

* `leaf_gigastream` (String) - leaf node stack gigastream alias
* `spine_box_id` (Number) - box-id of the spine nodes this spineLink is connected to
* `spine_gigastream` (String) - spine node gigastream alias
* `stack_link` (String) - Alias of the Stack Link this Spine Link is running over. Will be empty if corresponding StackLink is not yet created
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_spine_link.example {alias}
```
