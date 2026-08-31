---
page_title: "gigavuecore_get_aggregate_links Action - gigavuecore"
subcategory: ""
description: |-
  Link speed and count information between 2 set of nodes having different tag values for a given key
---

# gigavuecore_get_aggregate_links Action

Link speed and count information between 2 set of nodes having different tag values for a given key

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_get_aggregate_links" "example" {
  config {
    tags_filter = {
      endpoint1_tag_value = "example"
      endpoint2_tag_value = "example"
      link_type           = "example"
      tag_key             = "example"
    }
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `tags_filter` (Attributes, required) (see [below for nested schema](#nestedatt--tags_filter))

<a id="nestedatt--tags_filter"></a>
### Nested Schema for `tags_filter`

Required:

* `endpoint1_tag_value` (String)
* `endpoint2_tag_value` (String)
* `tag_key` (String)

Optional:

* `link_type` (String)

