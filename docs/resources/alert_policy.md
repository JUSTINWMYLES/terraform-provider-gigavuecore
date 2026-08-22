---
page_title: "gigavuecore_alert_policy Resource - gigavuecore"
subcategory: ""
description: |-
  Load Alert policy
---

# gigavuecore_alert_policy Resource

Load Alert policy

## Example Usage

```terraform
resource "gigavuecore_alert_policy" "example" {
  clear_condition = {}
  condition = {}
  description = null
  enabled = null
  metric = null
  policy_name = null
  resource_type = null
  resources = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `clear_condition` (Object({interval, threshold, type}), required)
  * `interval` (Object({duration, unit}), required)
    * `duration` (Number, required) - Averaging time duration
    * `unit` (String, required) - Unit of duration
  * `threshold` (Object({severity, threshold}), optional)
    * `severity` (String, required) - Severity of the alert
    * `threshold` (Number, required) - Threshold percentage value
  * `type` (String, required) - Alert clear type
* `condition` (Object({interval, thresholds}), required)
  * `interval` (Object({duration, unit}), required)
    * `duration` (Number, required) - Averaging time duration
    * `unit` (String, required) - Unit of duration
  * `thresholds` (List(Object({severity, threshold})), required) - Alert threshold definitions
* `description` (String, optional) - Description of the alert policy
* `enabled` (Bool, optional) - Status of the alert policy
* `metric` (String, required) - Metric name for which the alert policy is created
* `policy_name` (String, required) - Name of the alert policy, like an alias
* `resource_type` (String, optional) - Type of the resource, only tunnelMonitoring is supported in 6.1
* `resources` (Dynamic, required) - Resources associated to alert policy. Based on source type, the resources structure varies

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - Description of the alert policy
* `enabled` (Bool, computed) - Status of the alert policy
* `resource_type` (String, computed) - Type of the resource, only tunnelMonitoring is supported in 6.1

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_alert_policy.example {policy_name}
```
