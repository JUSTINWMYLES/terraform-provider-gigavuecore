---
page_title: "gigavuecore_send_email Data Source - gigavuecore"
subcategory: ""
description: |-
  Emails the usage report as encoded JSON for period containing date in format YYYYMMDD
---

# gigavuecore_send_email Data Source

Emails the usage report as encoded JSON for period containing date in format YYYYMMDD

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_send_email" "example" {
  date = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `date` (Number, required) - Date of processed volumes to retrieve in format YYYYMMDD


