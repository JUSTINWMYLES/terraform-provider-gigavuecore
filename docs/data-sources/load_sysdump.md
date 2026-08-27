---
page_title: "gigavuecore_load_sysdump Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Available Sysdump Filenames
---

# gigavuecore_load_sysdump Data Source

Load Available Sysdump Filenames

## Example Usage

```terraform
data "gigavuecore_load_sysdump" "example" {
  box_id     = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID range from 1 to 64(inclusive). all is applicable
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) - List of available sysdump files (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (String) - The chassis Box ID value
* `hostname_sysdump_status` (String) - The hostname of device for which sysdump generation status is sent
* `sysdump_files_per_box` (Attributes List) - List of available sysdump files (see [below for nested schema](#nestedatt--items--sysdump_files_per_box))
<a id="nestedatt--items--sysdump_files_per_box"></a>
### Nested Schema for `items.sysdump_files_per_box`

Read-Only:

* `filename` (String)
* `size` (Number) - Size of file in bytes
* `timestamp` (String) - File creation date and time in ISO 8601 format

