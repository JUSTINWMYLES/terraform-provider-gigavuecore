---
page_title: "gigavuecore_load_audit_log_entries List Resource - gigavuecore"
subcategory: ""
description: |-
  Load User Action Audit Log Entries
---

# gigavuecore_load_audit_log_entries List Resource

Load User Action Audit Log Entries

## Example Usage

```terraform
list "gigavuecore_load_audit_log_entries" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    end_time   = "example"
    operation  = "example"
    outcome    = "example"
    page       = "example"
    sort       = "example"
    start_time = "example"
    target     = "example"
    username   = "example"
  }
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


### Identity Attributes

The following identity attributes are exported for each matching result:

* `entry_id` (String, computed)


