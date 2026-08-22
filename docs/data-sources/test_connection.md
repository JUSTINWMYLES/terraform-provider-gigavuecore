---
page_title: "gigavuecore_test_connection Data Source - gigavuecore"
subcategory: ""
description: |-
  Tests the connectivity and auth is working for the api key
---

# gigavuecore_test_connection Data Source

Tests the connectivity and auth is working for the api key

## Example Usage

```terraform
data "gigavuecore_test_connection" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `message` (String, computed) - Copilot Connection status message
* `status` (String, computed) - Copilot Connection status code reason string
* `status_code` (Number, computed)

