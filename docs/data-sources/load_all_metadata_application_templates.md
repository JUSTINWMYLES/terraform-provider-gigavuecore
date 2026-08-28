---
page_title: "gigavuecore_load_all_metadata_application_templates Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Metadata Application Templates
---

# gigavuecore_load_all_metadata_application_templates Data Source

Load All Metadata Application Templates

## Example Usage

```terraform
data "gigavuecore_load_all_metadata_application_templates" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--items--attributes))
* `description` (String) - User friendly application name
* `family` (String) - application category name
* `name` (String) - application name
<a id="nestedatt--items--attributes"></a>
### Nested Schema for `items.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value

