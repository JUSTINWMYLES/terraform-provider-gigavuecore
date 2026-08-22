---
page_title: "gigavuecore_load_snmp_throttle_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Snmp Throttle config
---

# gigavuecore_load_snmp_throttle_config Data Source

Load Snmp Throttle config

## Example Usage

```terraform
data "gigavuecore_load_snmp_throttle_config" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({interval, notify_set, report_threshold, throttle_events})), computed) - list of SNMP Throttle Config Details on the node

