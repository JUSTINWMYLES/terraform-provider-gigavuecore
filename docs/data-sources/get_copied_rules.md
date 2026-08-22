---
page_title: "gigavuecore_get_copied_rules Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve previously copied rules from the user-specific clipboard filtered by rule category and type. Returns the complete set of copied rules matching the specified category and type.
---

# gigavuecore_get_copied_rules Data Source

Retrieve previously copied rules from the user-specific clipboard filtered by rule category and type. Returns the complete set of copied rules matching the specified category and type.

## Example Usage

```terraform
data "gigavuecore_get_copied_rules" "example" {
  rule_category = null
  rule_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `rule_category` (String, required) - Rule category (SOURCE or APPLICATION)
* `rule_type` (String, required) - Rule type (MapSubType enum value)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `copied_rules` (Object({created_time, policy_id, rule_category, rule_set, rule_type, username}), computed)
  * `created_time` (Number, computed) - Timestamp when rules were copied
  * `policy_id` (String, computed) - MongoDB document ID
  * `rule_category` (String, computed) - Category of rules to copy/paste
  * `rule_set` (List(Object({application_rules, flow_alias, policy_alias, policy_id, source_and_rule_alias, source_rules, sub_flow_alias})), computed) - Set of copied rules
  * `rule_type` (String, computed) - Type of application rule
  * `username` (String, computed) - Username of the user who copied the rules

