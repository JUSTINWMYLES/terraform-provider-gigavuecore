---
page_title: "gigavuecore_load_gs_dump Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Available Gigasmart dump Filenames
---

# gigavuecore_load_gs_dump Data Source

Load Available Gigasmart dump Filenames

## Example Usage

```terraform
data "gigavuecore_load_gs_dump" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({completed_list, eport_list, failed_list, filename, gs_exec_status, hostname, size, timestamp})), computed) - List of available Gigasmart Dump files

