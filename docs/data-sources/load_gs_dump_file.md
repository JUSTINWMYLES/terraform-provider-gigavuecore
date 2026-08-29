---
page_title: "gigavuecore_load_gs_dump_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Download Gigasmart dump File
---

# gigavuecore_load_gs_dump_file Data Source

Download Gigasmart dump File

## Example Usage

```terraform
data "gigavuecore_load_gs_dump_file" "example" {
  cluster_id = "example"
  filename   = "example"
  hostname   = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster Id
* `filename` (String, required) - Gigasmart dump filename to download
* `hostname` (String, required) - Target Host Name


