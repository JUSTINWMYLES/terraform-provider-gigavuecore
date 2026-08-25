---
page_title: "gigavuecore_load_alert_policies Data Source - gigavuecore"
subcategory: ""
description: |-
  Alert policy listing
---

# gigavuecore_load_alert_policies Data Source

Alert policy listing

## Example Usage

```terraform
data "gigavuecore_load_alert_policies" "example" {
  enabled       = null
  page          = null
  policy_name   = null
  resource_type = null
  sort          = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `enabled` (Boolean, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `policy_name` (String, optional) - Name of the alert policy
* `resource_type` (String, optional)
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alert_policies` (Attributes List, computed) - All Alert Policy Configurations (see [below for nested schema](#nestedatt--alert_policies))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--alert_policies"></a>
### Nested Schema for `alert_policies`

Read-Only:

* `clear_condition` (Attributes) (see [below for nested schema](#nestedatt--alert_policies--clear_condition))
* `condition` (Attributes) (see [below for nested schema](#nestedatt--alert_policies--condition))
* `description` (String) - Description of the alert policy
* `enabled` (Boolean) - Status of the alert policy
* `metric` (String) - Metric name for which the alert policy is created
* `policy_name` (String) - Name of the alert policy, like an alias
* `resource_type` (String) - Type of the resource, only tunnelMonitoring is supported in 6.1
* `resources` (Dynamic) - Resources associated to alert policy. Based on source type, the resources structure varies
<a id="nestedatt--alert_policies--clear_condition"></a>
### Nested Schema for `alert_policies.clear_condition`

Read-Only:

* `interval` (Attributes) (see [below for nested schema](#nestedatt--alert_policies--clear_condition--interval))
* `threshold` (Attributes) (see [below for nested schema](#nestedatt--alert_policies--clear_condition--threshold))
* `type` (String) - Alert clear type
<a id="nestedatt--alert_policies--clear_condition--interval"></a>
### Nested Schema for `alert_policies.clear_condition.interval`

Read-Only:

* `duration` (Number) - Averaging time duration
* `unit` (String) - Unit of duration
<a id="nestedatt--alert_policies--clear_condition--threshold"></a>
### Nested Schema for `alert_policies.clear_condition.threshold`

Read-Only:

* `severity` (String) - Severity of the alert
* `threshold` (Number) - Threshold percentage value
<a id="nestedatt--alert_policies--condition"></a>
### Nested Schema for `alert_policies.condition`

Read-Only:

* `interval` (Attributes) (see [below for nested schema](#nestedatt--alert_policies--condition--interval))
* `thresholds` (Attributes List) - Alert threshold definitions (see [below for nested schema](#nestedatt--alert_policies--condition--thresholds))
<a id="nestedatt--alert_policies--condition--interval"></a>
### Nested Schema for `alert_policies.condition.interval`

Read-Only:

* `duration` (Number) - Averaging time duration
* `unit` (String) - Unit of duration
<a id="nestedatt--alert_policies--condition--thresholds"></a>
### Nested Schema for `alert_policies.condition.thresholds`

Read-Only:

* `severity` (String) - Severity of the alert
* `threshold` (Number) - Threshold percentage value
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

