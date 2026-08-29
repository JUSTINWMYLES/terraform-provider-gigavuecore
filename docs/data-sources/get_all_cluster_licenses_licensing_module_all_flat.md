---
page_title: "gigavuecore_get_all_cluster_licenses_licensing_module_all_flat Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all clusters with all the chassis and cards information and licenses installed in List View
---

# gigavuecore_get_all_cluster_licenses_licensing_module_all_flat Data Source

Get all clusters with all the chassis and cards information and licenses installed in List View

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_licenses_licensing_module_all_flat" "example" {
  cluster_name  = "example"
  end_date      = 0
  feature       = [ "example" ]
  host_name     = "example"
  page          = "example"
  serial_number = "example"
  slot_id       = "example"
  sort          = "example"
  start_date    = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_name` (String, optional) - Get all clusters by Cluster Name
* `end_date` (Number, optional) - Get all clusters by EndDate
* `feature` (List of String, optional) - Get all clusters by Feature
* `host_name` (String, optional) - Get all clusters by Host Name
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `serial_number` (String, optional) - Get all clusters by SerialNumber
* `slot_id` (String, optional) - Get all clusters by SlotId
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `start_date` (Number, optional) - Get all clusters by StartDate

### Attributes

In addition to all arguments above, the following attributes are exported:

* `box_id` (String, computed)
* `chassis_serial_number` (String, computed)
* `cluster_name` (String, computed) - Get all clusters by Cluster Name
* `device_sw_version` (String, computed)
* `end_date` (Number, computed) - Get all clusters by EndDate
* `feature` (List of String, computed) - Get all clusters by Feature
* `floated` (String, computed)
* `grace_period` (Number, computed)
* `host_name` (String, computed) - Get all clusters by Host Name
* `hw_type` (String, computed)
* `license_status` (String, computed)
* `license_type` (String, computed)
* `model` (String, computed)
* `node_id` (String, computed)
* `notif_period` (Number, computed)
* `serial_number` (String, computed) - Get all clusters by SerialNumber
* `slot_id` (String, computed) - Get all clusters by SlotId
* `start_date` (Number, computed) - Get all clusters by StartDate


