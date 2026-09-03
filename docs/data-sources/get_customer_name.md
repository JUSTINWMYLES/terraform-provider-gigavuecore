---
page_title: "gigavuecore_get_customer_name Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns the name of the customer that owns this FM
---

# gigavuecore_get_customer_name Data Source

Returns the name of the customer that owns this FM

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_get_customer_name" "example" {
}
```
