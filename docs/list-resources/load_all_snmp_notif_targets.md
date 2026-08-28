---
page_title: "gigavuecore_load_all_snmp_notif_targets List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all SNMP Notification Targets
---

# gigavuecore_load_all_snmp_notif_targets List Resource

Load all SNMP Notification Targets

## Example Usage

```terraform
list "gigavuecore_load_all_snmp_notif_targets" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
    page       = "example"
    sort       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `notif_target_address` (String, computed)


