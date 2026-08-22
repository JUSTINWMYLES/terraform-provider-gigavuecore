---
page_title: "gigavuecore_create_api_token Data Source - gigavuecore"
subcategory: ""
description: |-
  Get FM API rate limit configuration.
---

# gigavuecore_create_api_token Data Source

Get FM API rate limit configuration.

## Example Usage

```terraform
data "gigavuecore_create_api_token" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `request_limit_count` (Number, computed) - FM API Request Limit Count
* `request_limit_duration` (Number, computed) - FM API Request Limit Duration in seconds

