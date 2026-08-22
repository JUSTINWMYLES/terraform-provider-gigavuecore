---
page_title: "gigavuecore_configure_expiry_alert_enable Action - gigavuecore"
subcategory: ""
description: |-
  Configure license expiry alert enable
---

# gigavuecore_configure_expiry_alert_enable Action

Configure license expiry alert enable

## Example Usage

```terraform
action "gigavuecore_configure_expiry_alert_enable" "example" {
  config {
    enable = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `enable` (Bool, required) - the enable value
