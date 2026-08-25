---
page_title: "gigavuecore_create_api_token_fm_system_api_rate_limiting Action - gigavuecore"
subcategory: ""
description: |-
  Modify FM API rate limit configuration.
---

# gigavuecore_create_api_token_fm_system_api_rate_limiting Action

Modify FM API rate limit configuration.

## Example Usage

```terraform
action "gigavuecore_create_api_token_fm_system_api_rate_limiting" "example" {
  config {
    request_limit_count    = 1
    request_limit_duration = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `request_limit_count` (Number, required) - FM API Request Limit Count
* `request_limit_duration` (Number, required) - FM API Request Limit Duration in seconds


