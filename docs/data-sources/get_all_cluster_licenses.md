---
page_title: "gigavuecore_get_all_cluster_licenses Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all clusters with all the chassis and cards information and licenses installed in  Hierarchical View
---

# gigavuecore_get_all_cluster_licenses Data Source

Get all clusters with all the chassis and cards information and licenses installed in  Hierarchical View

## Example Usage

```terraform
data "gigavuecore_get_all_cluster_licenses" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({chassis_modules, cluster_name, hash_code})), computed)

