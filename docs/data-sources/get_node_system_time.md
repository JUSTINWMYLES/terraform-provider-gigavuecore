---
page_title: "gigavuecore_get_node_system_time Data Source - gigavuecore"
subcategory: ""
description: |-
  Get current system time
---

# gigavuecore_get_node_system_time Data Source

Get current system time

## Example Usage

```terraform
data "gigavuecore_get_node_system_time" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `time` (String, computed) - In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `time_zone` (String, computed) - TZ in the form Area/Location or Etc/GMT, e.g. America/Los\_Angeles, Etc/GMT-8
* `utc_offset` (Number, computed) - Read-only. UTC Offset in seconds


