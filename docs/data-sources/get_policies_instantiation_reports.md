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
  page = null
  policy_id = null
  since = null
  sort = null
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

* `av_policy_trigger_reports` (List(Object({action_status, condition_status, outcome, policy_id, policy_name, trigger_time})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

