---
page_title: "gigavuecore_get_policies_instantiation_reports Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Active Visibility Trigger Reports
---

# gigavuecore_get_policies_instantiation_reports Data Source

Get Active Visibility Trigger Reports

## Example Usage

```terraform
data "gigavuecore_get_policies_instantiation_reports" "example" {
  page       = null
  policy_id  = null
  since      = null
  sort       = null
  time_range = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `policy_id` (String, optional) - Target Policy Id. If left out, reports for all AV Policies are returned
* `since` (String, optional) - Reference timestamps to include trigger reports from. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'. Mutually exclusive with 'timeRange' query parameter
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `time_range` (String, optional) - Time window for which to include trigger reports. Mutually exclusive with 'since' query paramater

### Attributes

In addition to all arguments above, the following attributes are exported:

* `av_policy_trigger_reports` (Attributes List, computed) (see [below for nested schema](#nestedatt--av_policy_trigger_reports))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--av_policy_trigger_reports"></a>
### Nested Schema for `av_policy_trigger_reports`

Read-Only:

* `action_status` (Attributes List) - Execution report for all associated Policy Actions. On Policy 'success', could be left out since success of every action is implied. (see [below for nested schema](#nestedatt--av_policy_trigger_reports--action_status))
* `condition_status` (Attributes List) - Triggering Conditions report. If none of the conditions are parameterizable, this status field can be omitted (see [below for nested schema](#nestedatt--av_policy_trigger_reports--condition_status))
* `outcome` (String) - 'success' is reported when all actions completed successfully, otherwise 'failure' is reported
* `policy_id` (String) - reference to the triggering Policy
* `policy_name` (String) - triggering Policy name. Included for readability to provide Policy context
* `trigger_time` (String) - policy trigger time. In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
<a id="nestedatt--av_policy_trigger_reports--action_status"></a>
### Nested Schema for `av_policy_trigger_reports.action_status`

Read-Only:

* `action_id` (String) - Action ID in the parent Policy
* `action_type` (String) - Action Type (Source Action Template). Included for readability to provide Action context
* `failure_reasons` (List of String) - List of the reasons that resulted in the failed Action execution. Empty or omitted on success
* `outcome` (String)
<a id="nestedatt--av_policy_trigger_reports--condition_status"></a>
### Nested Schema for `av_policy_trigger_reports.condition_status`

Read-Only:

* `condition_id` (String) - Condition ID in the parent Policy
* `params` (Attributes List) - key/value criteria parameters under which this Condition was satisfied for a given Policy triggering instance. This map should have an entry for each Parameter required for the corresponding condition type. If corresponding Condition does not take parameters, this field can be omitted (see [below for nested schema](#nestedatt--av_policy_trigger_reports--condition_status--params))
<a id="nestedatt--av_policy_trigger_reports--condition_status--params"></a>
### Nested Schema for `av_policy_trigger_reports.condition_status.params`

Read-Only:

* `key` (String)
* `value` (String)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

