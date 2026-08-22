---
page_title: "gigavuecore_get_gtp_backup_files Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Gtp backup files
---

# gigavuecore_get_gtp_backup_files Data Source

Load all Gtp backup files

## Example Usage

```terraform
data "gigavuecore_get_gtp_backup_files" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({cluster_id, filename, size, timestamp})), computed)

