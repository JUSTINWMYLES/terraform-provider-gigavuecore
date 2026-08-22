---
page_title: "gigavuecore_get_avisi_condition_templ Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Active Visibility Condition Templates
---

# gigavuecore_get_avisi_condition_templ Data Source

Get Active Visibility Condition Templates

## Example Usage

```terraform
data "gigavuecore_get_avisi_condition_templ" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({description, name, template_desc})), computed)

