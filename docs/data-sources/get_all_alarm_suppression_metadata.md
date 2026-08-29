---
page_title: "gigavuecore_get_all_alarm_suppression_metadata Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Alarm Supppression Rules
---

# gigavuecore_get_all_alarm_suppression_metadata Data Source

Get All Alarm Supppression Rules

## Example Usage

```terraform
data "gigavuecore_get_all_alarm_suppression_metadata" "example" {
  alarm_types     = "example"
  alias           = "example"
  cluster_id      = "example"
  hostname        = "example"
  page            = "example"
  resource_id     = "example"
  selected_reason = "example"
  sort            = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alarm_types` (String, optional) - comma-separated list of Alarm Type's to filter
* `alias` (String, optional)
* `cluster_id` (String, optional)
* `hostname` (String, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident
* `resource_id` (String, optional)
* `selected_reason` (String, optional)
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alarm_types` (List of String)
* `alias` (String) - Alias
* `cluster_id` (String) - Cluster/Node Id
* `created_by` (String) - Alarm suppression rule created by FM User
* `created_ts` (String) - Alarm suppression rule creation timestamp in ISO 8601 format
* `enable` (Boolean) - FM Alarm Suppression rule state
* `expiry_time` (Number) - Expiry interval for suppression rule
* `expiry_time_unit` (String) - Expiry interval unit
* `expiry_ts` (String) - Expiry Timestamp(UTC) for suppression rule
* `hostname` (String) - FM Hostname
* `resource_id` (String) - Id of the resource to be suppressed
* `resource_type` (String)
* `selected_reason` (String) - Suppression Reason
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--items--tags))
* `updated_by` (String) - Alarm suppression rule updated by FM User
* `updated_ts` (String) - Alarm suppression rule updated timestamp in ISO 8601 format

<a id="nestedatt--items--tags"></a>
### Nested Schema for `items.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

