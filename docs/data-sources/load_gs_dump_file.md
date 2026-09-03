---
page_title: "gigavuecore_load_gs_dump_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Download Gigasmart dump File
---

# gigavuecore_load_gs_dump_file Data Source

Download Gigasmart dump File

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

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


