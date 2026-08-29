---
page_title: "gigavuecore_load_aliases_meta_data Action - gigavuecore"
subcategory: ""
description: |-
  Get meta data details for given list of resources
---

# gigavuecore_load_aliases_meta_data Action

Get meta data details for given list of resources

## Example Usage

```terraform
action "gigavuecore_load_aliases_meta_data" "example" {
  config {
    specs = [{
      cluster_id   = "example"
      resource_ids = [ "example" ]
      type         = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `specs` (Attributes List, required) (see [below for nested schema](#nestedatt--specs))

<a id="nestedatt--specs"></a>
### Nested Schema for `specs`

Required:

* `cluster_id` (String)
* `resource_ids` (List of String)
* `type` (String)

