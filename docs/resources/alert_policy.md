---
page_title: "gigavuecore_alert_policy Resource - gigavuecore"
subcategory: ""
description: |-
  Creation of alert policy
---

# gigavuecore_alert_policy Resource

Creation of alert policy

## Example Usage

```terraform
resource "gigavuecore_alert_policy" "example" {
  clear_condition = {
    interval = {
      duration = 0
      unit     = "MINUTES"
    }
    threshold = {
      severity  = "Info"
      threshold = 1.0
    }
    type = "TIME_BASED"
  }
  condition = {
    interval = {
      duration = 0
      unit     = "MINUTES"
    }
    thresholds = [{
      severity  = "Info"
      threshold = 1.0
    }]
  }
  description   = "example"
  enabled       = true
  metric        = "TUNNEL_TRAFFIC_MONITORING"
  policy_name   = "example"
  resource_type = "tunnelMonitoring"
  resources     = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `clear_condition` (Attributes, required) - Alert clear condition definition (see [below for nested schema](#nestedatt--clear_condition))
* `condition` (Attributes, required) - Alert condition definition (see [below for nested schema](#nestedatt--condition))
* `description` (String, optional) - Description of the alert policy
* `enabled` (Boolean, optional) - Status of the alert policy
* `metric` (String, required) - Metric name for which the alert policy is created
* `policy_name` (String, required) - Name of the alert policy, like an alias
* `resource_type` (String, optional) - Type of the resource, only tunnelMonitoring is supported in 6.1
* `resources` (Dynamic, required) - Resources associated to alert policy. Based on source type, the resources structure varies

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--clear_condition"></a>
### Nested Schema for `clear_condition`

Required:

* `interval` (Attributes) - Alert clear interval config (see [below for nested schema](#nestedatt--clear_condition--interval))
* `type` (String) - Alert clear type

Optional:

* `threshold` (Attributes) - Alert clear threshold definition (see [below for nested schema](#nestedatt--clear_condition--threshold))

<a id="nestedatt--clear_condition--interval"></a>
### Nested Schema for `clear_condition.interval`

Required:

* `duration` (Number) - Averaging time duration
* `unit` (String) - Unit of duration

<a id="nestedatt--clear_condition--threshold"></a>
### Nested Schema for `clear_condition.threshold`

Required:

* `severity` (String) - Severity of the alert
* `threshold` (Number) - Threshold percentage value

<a id="nestedatt--condition"></a>
### Nested Schema for `condition`

Required:

* `interval` (Attributes) - Alert interval config (see [below for nested schema](#nestedatt--condition--interval))
* `thresholds` (Attributes List) - Alert threshold definitions (see [below for nested schema](#nestedatt--condition--thresholds))

<a id="nestedatt--condition--interval"></a>
### Nested Schema for `condition.interval`

Required:

* `duration` (Number) - Averaging time duration
* `unit` (String) - Unit of duration

<a id="nestedatt--condition--thresholds"></a>
### Nested Schema for `condition.thresholds`

Required:

* `severity` (String) - Severity of the alert
* `threshold` (Number) - Threshold percentage value
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_alert_policy.example {policy_name}
```
