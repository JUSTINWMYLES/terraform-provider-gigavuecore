---
page_title: "gigavuecore_get_audit_log_entry_by_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Find Audit Log Entry by ID
---

# gigavuecore_get_audit_log_entry_by_id Data Source

Find Audit Log Entry by ID

## Example Usage

```terraform
data "gigavuecore_get_audit_log_entry_by_id" "example" {
  entry_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `entry_id` (String, required) - ID of the target Audit Log Entry

### Attributes

In addition to all arguments above, the following attributes are exported:

* `description` (String, computed) - Event description. In case of failure, holds error description
* `operation` (String, computed) - Description of the user action
* `outcome` (String, computed) - Scope of the event
* `target` (String, computed) - Action target. Device ID for node configuration actions. Component ID for FM-based service configuration
* `ts` (String, computed) - Event timestamp in ISO 8601 format
* `ts_utc` (Number, computed) - Timestamp in UTC milliseconds
* `update_details` (Dynamic, computed) - Optional extension to hold the update delta details
* `username` (String, computed) - Action target. Device ID for node configuration actions. Component ID for FM-based service configuration

