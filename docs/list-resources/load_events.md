---
page_title: "gigavuecore_load_events List Resource - gigavuecore"
subcategory: ""
description: |-
  Load Events
---

# gigavuecore_load_events List Resource

Load Events

## Example Usage

```terraform
list "gigavuecore_load_events" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    end_time      = "example"
    page          = "example"
    resource_id   = "example"
    resource_type = "example"
    scope         = "example"
    severity      = "example"
    sort          = "example"
    source        = "example"
    start_time    = "example"
    type          = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `end_time` (String, optional) - End Time to filter by. In ISO 8601 format
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `resource_id` (String, optional) - Affected Entity to filter by
* `resource_type` (String, optional) - Affected Entity Type to filter by
* `scope` (String, optional) - Event Scope to filter by
* `severity` (String, optional) - Event Scope to filter by
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `source` (String, optional) - Event Source identifier to filter by
* `start_time` (String, optional) - Start Time to filter by. In ISO 8601 format
* `type` (String, optional) - Event Type to filter by


### Identity Attributes

The following identity attributes are exported for each matching result:

* `event_id` (String, computed)


