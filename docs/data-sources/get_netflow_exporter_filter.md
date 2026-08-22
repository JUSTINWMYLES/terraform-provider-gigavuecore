---
page_title: "gigavuecore_get_netflow_exporter_filter Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Netflow Exporter Filter
---

# gigavuecore_get_netflow_exporter_filter Data Source

Get Netflow Exporter Filter

## Example Usage

```terraform
data "gigavuecore_get_netflow_exporter_filter" "example" {
  alias = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Netflow Exporter
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({pass_rules})), computed)

