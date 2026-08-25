---
page_title: "gigavuecore_send_email2 Data Source - gigavuecore"
subcategory: ""
description: |-
  Emails the usage report in PDF format for period containing date in format YYYYMMDD
---

# gigavuecore_send_email2 Data Source

Emails the usage report in PDF format for period containing date in format YYYYMMDD

## Example Usage

```terraform
data "gigavuecore_send_email2" "example" {
  date = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `date` (Number, required) - Date of processed volumes to retrieve in format YYYYMMDD


