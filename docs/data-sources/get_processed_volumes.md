---
page_title: "gigavuecore_get_processed_volumes Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the processed volumes of the specified date which include application tier usage and application usage in bytes
---

# gigavuecore_get_processed_volumes Data Source

Gives the processed volumes of the specified date which include application tier usage and application usage in bytes

## Example Usage

```terraform
data "gigavuecore_get_processed_volumes" "example" {
  date = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `date` (Number, required) - Date of processed volumes to retrieve in format YYYYMMDD

### Attributes

In addition to all arguments above, the following attributes are exported:

* `any` (Map of Number, computed)
* `any0` (Number, computed)
* `app` (Map of Number, computed)


