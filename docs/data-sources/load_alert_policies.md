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
  enabled = null
  page = null
  policy_name = null
  resource_type = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `enabled` (Bool, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `policy_name` (String, optional) - Name of the alert policy
* `resource_type` (String, optional)
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alert_policies` (List(Object({clear_condition, condition, description, enabled, metric, policy_name, resource_type, resources})), computed) - All Alert Policy Configurations
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

