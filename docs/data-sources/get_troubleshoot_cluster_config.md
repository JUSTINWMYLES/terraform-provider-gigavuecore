---
page_title: "gigavuecore_get_troubleshoot_cluster_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve troubleshoot cluster configuration for a traffic flow
---

# gigavuecore_get_troubleshoot_cluster_config Data Source

Retrieve troubleshoot cluster configuration for a traffic flow

## Example Usage

```terraform
data "gigavuecore_get_troubleshoot_cluster_config" "example" {
  alias          = "example"
  cluster_id     = "example"
  connected_maps = true
  meta_data      = true
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias
* `cluster_id` (String, required) - Target Cluster ID
* `connected_maps` (Boolean, optional) - Include connected maps in the response
* `meta_data` (Boolean, optional) - Include cluster config metadata in the response

### Attributes

In addition to all arguments above, the following attributes are exported:

* `value` (Dynamic, computed)


