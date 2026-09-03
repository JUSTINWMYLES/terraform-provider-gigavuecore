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
  end_time   = "example"
  operation  = "example"
  outcome    = "example"
  page       = "example"
  sort       = "example"
  start_time = "example"
  target     = "example"
  username   = "example"
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `description` (String) - Event description. In case of failure, holds error description
* `operation` (String) - Description of the user action
* `outcome` (String) - Scope of the event
* `target` (String) - Action target. Device ID for node configuration actions. Component ID for FM-based service configuration
* `ts` (String) - Event timestamp in ISO 8601 format
* `ts_utc` (Number) - Timestamp in UTC milliseconds
* `update_details` (Dynamic) - Optional extension to hold the update delta details
* `username` (String) - Action target. Device ID for node configuration actions. Component ID for FM-based service configuration

