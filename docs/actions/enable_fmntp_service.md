---
page_title: "gigavuecore_enable_fmntp_service Action - gigavuecore"
subcategory: ""
description: |-
  Enable/Disable NTP Service in FM
---

# gigavuecore_enable_fmntp_service Action

Enable/Disable NTP Service in FM

## Example Usage

```terraform
action "gigavuecore_enable_fmntp_service" "example" {
  config {
    enabled = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `enabled` (Bool, required) - NTP service enable status
