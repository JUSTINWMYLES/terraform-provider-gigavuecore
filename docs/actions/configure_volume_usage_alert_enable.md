---
page_title: "gigavuecore_configure_volume_usage_alert_enable Action - gigavuecore"
subcategory: ""
description: |-
  Configure volume usage alert enable
---

# gigavuecore_configure_volume_usage_alert_enable Action

Configure volume usage alert enable

## Example Usage

```terraform
action "gigavuecore_configure_volume_usage_alert_enable" "example" {
  config {
    enable    = true
    threshold = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `enable` (Boolean, required) - the enable value
* `threshold` (Number, required) - the threshold value, positive (alert is triggered when usage reaches 'threshold' percentage of allowance)


