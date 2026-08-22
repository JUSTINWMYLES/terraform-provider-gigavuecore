---
page_title: "gigavuecore_get_avisi_policies Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Active Visibility Policies
---

# gigavuecore_get_avisi_policies Data Source

Get Active Visibility Policies

## Example Usage

```terraform
data "gigavuecore_get_avisi_policies" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({description, enabled, name, then_do, when_condition})), computed)

