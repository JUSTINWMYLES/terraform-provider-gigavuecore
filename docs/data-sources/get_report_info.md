---
page_title: "gigavuecore_get_report_info Data Source - gigavuecore"
subcategory: ""
description: |-
  Get metadata of Customer Deployed Assets Report
---

# gigavuecore_get_report_info Data Source

Get metadata of Customer Deployed Assets Report

## Example Usage

```terraform
data "gigavuecore_get_report_info" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `fm_user_token_entities` (Attributes List, computed) (see [below for nested schema](#nestedatt--fm_user_token_entities))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--fm_user_token_entities"></a>
### Nested Schema for `fm_user_token_entities`

Read-Only:

* `created_by` (String) - FM user who has created the CDA Audit Report
* `created_date` (String) - Created Time of the CDA Audit Report
* `file_name` (String) - Name of the CDA Audit Report
* `file_size` (String) - CDA Audit Report file size

