---
page_title: "gigavuecore_load_licensing_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Licensing Config
---

# gigavuecore_load_licensing_config Data Source

Load Licensing Config

## Example Usage

```terraform
data "gigavuecore_load_licensing_config" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `expired_licenses_exist` (Boolean, computed)
* `expiry_alerts_last_sent_on` (Number, computed)
* `license_expiry_alert_enabled` (Boolean, computed)
* `mandatory_upcoming_renewals_receiver` (String, computed)
* `mandatory_vbl_report_receiver` (String, computed)
* `months_per_period` (Number, computed)
* `other_upcoming_renewals_receivers` (List of String, computed)
* `other_vbl_report_receivers` (List of String, computed)
* `volume_usage_alert_enabled` (Boolean, computed)
* `volume_usage_alert_threshold` (Number, computed)


