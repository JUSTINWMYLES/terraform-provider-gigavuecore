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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `created_by` (String) - FM user who has created the CDA Audit Report
* `created_date` (String) - Created Time of the CDA Audit Report
* `file_name` (String) - Name of the CDA Audit Report
* `file_size` (String) - CDA Audit Report file size

