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
  box_id = null
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

* `items` (List(Object({box_id, hostname_sysdump_status, sysdump_files_per_box})), computed) - List of available sysdump files

