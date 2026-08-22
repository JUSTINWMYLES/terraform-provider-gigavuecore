---
page_title: "gigavuecore_get_flow_ops_report Data Source - gigavuecore"
subcategory: ""
description: |-
  Load FlowOps Report
---

# gigavuecore_get_flow_ops_report Data Source

Load FlowOps Report

## Example Usage

```terraform
data "gigavuecore_get_flow_ops_report" "example" {
  alias = null
  caller_id_pattern = null
  cluster_id = null
  type = null
  user_name_pattern = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `caller_id_pattern` (String, optional) - only valid if type is 'flowSip'
* `cluster_id` (String, required) - Target Cluster ID
* `type` (String, required) - Flowops report type
* `user_name_pattern` (String, optional) - only valid if type is 'flowDiameterS6a'

### Attributes

In addition to all arguments above, the following attributes are exported:

* `report` (String, computed)

