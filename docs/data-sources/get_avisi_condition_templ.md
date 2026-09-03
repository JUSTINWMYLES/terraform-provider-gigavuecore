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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `description` (String)
* `name` (String) - User-friendly AV Condition template name
* `template_desc` (String) - User-friendly Condition template with associated named placeholders

