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

* `items` (Attributes List, computed) - List of available Gigasmart Dump files (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `completed_list` (String) - Eport list for which GigaSmart Dump is completed
* `eport_list` (String) - Complete list of Eports for GigaSmart dump generation
* `failed_list` (String) - Eport list for which GigaSmart Dump failed
* `filename` (String) - Filename of the Gigasmart dump file
* `gs_exec_status` (String) - Execution status of  GigaSmart Dump
* `hostname` (String) - Hostname of the device where Gigasmart dump resides
* `size` (Number) - Size of file in bytes
* `timestamp` (String) - File creation date and time in ISO 8601 format

