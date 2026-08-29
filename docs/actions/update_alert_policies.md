---
page_title: "gigavuecore_update_alert_policies Action - gigavuecore"
subcategory: ""
description: |-
  Update of list of given alert-policies
---

# gigavuecore_update_alert_policies Action

Update of list of given alert-policies

## Example Usage

```terraform
action "gigavuecore_update_alert_policies" "example" {
  config {
    context = {
      page_no     = 0
      page_size   = 0
      sort        = [ "example" ]
      total_items = 0
    }
    policies = [{
      clear_condition = {
        interval = {
          duration = 0
          unit     = "example"
        }
        threshold = {
          severity  = "example"
          threshold = 1.0
        }
        type = "example"
      }
      condition = {
        interval = {
          duration = 0
          unit     = "example"
        }
        thresholds = [{
          severity  = "example"
          threshold = 1.0
        }]
      }
      description     = "example"
      enabled         = true
      metric          = "example"
      new_policy_name = "example"
      policy_name     = "example"
      resource_type   = "example"
      resources       = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `context` (Attributes, optional) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `policies` (Attributes List, required) - List of alert policies to be updated (see [below for nested schema](#nestedatt--policies))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Required:

* `total_items` (Number) - total number of items in the queried entity type

Optional:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order

<a id="nestedatt--policies"></a>
### Nested Schema for `policies`

Required:

* `clear_condition` (Attributes) - Alert clear condition definition (see [below for nested schema](#nestedatt--policies--clear_condition))
* `condition` (Attributes) - Alert condition definition (see [below for nested schema](#nestedatt--policies--condition))
* `enabled` (Boolean) - Status of the alert policy
* `metric` (String) - Metric name for which the alert policy is created
* `policy_name` (String) - Name of the alert policy, like an alias
* `resources` (Dynamic) - Resources associated to alert policy. Based on source type, the resources structure varies

Optional:

* `description` (String) - Description of the alert policy
* `new_policy_name` (String) - New name for the alert policy, for changing the alert policy name this property needs to be provided
* `resource_type` (String) - Type of the resource, only tunnelMonitoring is supported in 6.3

<a id="nestedatt--policies--clear_condition"></a>
### Nested Schema for `policies.clear_condition`

Required:

* `interval` (Attributes) - Alert clear interval config (see [below for nested schema](#nestedatt--policies--clear_condition--interval))
* `type` (String) - Alert clear type

Optional:

* `threshold` (Attributes) - Alert clear threshold definition (see [below for nested schema](#nestedatt--policies--clear_condition--threshold))

<a id="nestedatt--policies--clear_condition--interval"></a>
### Nested Schema for `policies.clear_condition.interval`

Required:

* `duration` (Number) - Averaging time duration
* `unit` (String) - Unit of duration

<a id="nestedatt--policies--clear_condition--threshold"></a>
### Nested Schema for `policies.clear_condition.threshold`

Required:

* `severity` (String) - Severity of the alert
* `threshold` (Number) - Threshold percentage value

<a id="nestedatt--policies--condition"></a>
### Nested Schema for `policies.condition`

Required:

* `interval` (Attributes) - Alert interval config (see [below for nested schema](#nestedatt--policies--condition--interval))
* `thresholds` (Attributes List) - Alert threshold definitions (see [below for nested schema](#nestedatt--policies--condition--thresholds))

<a id="nestedatt--policies--condition--interval"></a>
### Nested Schema for `policies.condition.interval`

Required:

* `duration` (Number) - Averaging time duration
* `unit` (String) - Unit of duration

<a id="nestedatt--policies--condition--thresholds"></a>
### Nested Schema for `policies.condition.thresholds`

Required:

* `severity` (String) - Severity of the alert
* `threshold` (Number) - Threshold percentage value

