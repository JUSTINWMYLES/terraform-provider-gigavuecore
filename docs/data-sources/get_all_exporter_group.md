---
page_title: "gigavuecore_get_all_exporter_group Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Exporter Group
---

# gigavuecore_get_all_exporter_group Data Source

Get all Apps Exporter Group

## Example Usage

```terraform
data "gigavuecore_get_all_exporter_group" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `description` (String)
* `exporters` (List of String)

