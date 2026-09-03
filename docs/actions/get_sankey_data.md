---
page_title: "gigavuecore_get_sankey_data Action - gigavuecore"
subcategory: ""
description: |-
  Response for sankey representation
---

# gigavuecore_get_sankey_data Action

Response for sankey representation

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_get_sankey_data" "example" {
  config {
    nodes_filter = [{
      placement_tag_value = "example"
      sort = [{
        field = "example"
        order = "asc"
      }]
    }]
    tags_filter = {
      parent_tags = [{
        tag_key   = "example"
        tag_value = "example"
      }]
      placement_tag_alias = "example"
      tag_key             = "example"
      tag_value           = "example"
    }
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `nodes_filter` (Attributes List, optional) (see [below for nested schema](#nestedatt--nodes_filter))
* `tags_filter` (Attributes, required) (see [below for nested schema](#nestedatt--tags_filter))

<a id="nestedatt--nodes_filter"></a>
### Nested Schema for `nodes_filter`

Optional:

* `placement_tag_value` (String)
* `sort` (Attributes List) (see [below for nested schema](#nestedatt--nodes_filter--sort))

<a id="nestedatt--nodes_filter--sort"></a>
### Nested Schema for `nodes_filter.sort`

Optional:

* `field` (String)
* `order` (String)

<a id="nestedatt--tags_filter"></a>
### Nested Schema for `tags_filter`

Required:

* `placement_tag_alias` (String) - Configured placement tag alias
* `tag_key` (String) - Sankey tag key
* `tag_value` (String) - Sankey tag value

Optional:

* `parent_tags` (Attributes List) - Parent hierarchy information for the requested sankey (see [below for nested schema](#nestedatt--tags_filter--parent_tags))

<a id="nestedatt--tags_filter--parent_tags"></a>
### Nested Schema for `tags_filter.parent_tags`

Optional:

* `tag_key` (String)
* `tag_value` (String)

