---
page_title: "gigavuecore_refresh_all_cluster_licenses Data Source - gigavuecore"
subcategory: ""
description: |-
  Refresh all clusters with all the chassis and cards information and licenses installed therein
---

# gigavuecore_refresh_all_cluster_licenses Data Source

Refresh all clusters with all the chassis and cards information and licenses installed therein

## Example Usage

```terraform
data "gigavuecore_refresh_all_cluster_licenses" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(String), computed) - A list of db IDs of actually updated clusters, because of the refresh

