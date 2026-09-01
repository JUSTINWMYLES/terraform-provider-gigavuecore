---
page_title: "gigavuecore_get_cluster_state Data Source - gigavuecore"
subcategory: ""
description: |-
  Cluster state of FmHa
---

# gigavuecore_get_cluster_state Data Source

Cluster state of FmHa

## Example Usage

```terraform
data "gigavuecore_get_cluster_state" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `is_ready` (Boolean, computed) - Is HA cluster ready
* `status` (String, computed) - HA cluster status


