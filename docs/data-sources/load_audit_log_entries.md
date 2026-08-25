---
page_title: "gigavuecore_load_audit_log_entries Data Source - gigavuecore"
subcategory: ""
description: |-
  Load User Action Audit Log Entries
---

# gigavuecore_load_audit_log_entries Data Source

Load User Action Audit Log Entries

## Example Usage

```terraform
data "gigavuecore_load_audit_log_entries" "example" {
  end_time   = null
  operation  = null
  outcome    = null
  page       = null
  sort       = null
  start_time = null
  target     = null
  username   = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `end_time` (String, optional) - End Time to filter by. In ISO 8601 format
* `operation` (String, optional) - User Action type identifier to filter by
* `outcome` (String, optional) - User Action outcome type to filter by
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `start_time` (String, optional) - Start Time to filter by. In ISO 8601 format
* `target` (String, optional) - Action Target identifier to filter by
* `username` (String, optional) - Username to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `audit_log_entries` (Attributes List, computed) (see [below for nested schema](#nestedatt--audit_log_entries))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--audit_log_entries"></a>
### Nested Schema for `audit_log_entries`

Read-Only:

* `description` (String) - Event description. In case of failure, holds error description
* `operation` (String) - Description of the user action
* `outcome` (String) - Scope of the event
* `target` (String) - Action target. Device ID for node configuration actions. Component ID for FM-based service configuration
* `ts` (String) - Event timestamp in ISO 8601 format
* `ts_utc` (Number) - Timestamp in UTC milliseconds
* `update_details` (Dynamic) - Optional extension to hold the update delta details
* `username` (String) - Action target. Device ID for node configuration actions. Component ID for FM-based service configuration
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

