---
page_title: "gigavuecore_update_node_system_time Action - gigavuecore"
subcategory: ""
description: |-
  Update Node System Time configuration
---

# gigavuecore_update_node_system_time Action

Update Node System Time configuration

## Example Usage

```terraform
action "gigavuecore_update_node_system_time" "example" {
  config {
    cluster_id = "example"
    time = "example"
    time_zone = "example"
    utc_offset = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `time` (String, optional) - In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'
* `time_zone` (String, optional) - TZ in the form Area/Location or Etc/GMT, e.g. America/Los\_Angeles, Etc/GMT-8
* `utc_offset` (Number, optional) - Read-only. UTC Offset in seconds
