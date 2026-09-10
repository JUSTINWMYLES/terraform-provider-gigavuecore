---
page_title: "gigavuecore_get_fabric_resource_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Get current fabric resource allocation policy configuration
---

# gigavuecore_get_fabric_resource_config Data Source

Get current fabric resource allocation policy configuration

## Example Usage

```terraform
data "gigavuecore_get_fabric_resource_config" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `mode` (String, computed) - \['SHARE' or 'NOT\_SHARE'\]: Resource sharing mode. SHARE: same resource can be shared by different fabric maps; NOT\_SHARE: not shared. Default: NOT\_SHARE.
* `scope` (String, computed) - \['GLOBAL'\]: Scope of resource pool. GLOBAL: only one global resource pool.
* `type` (String, computed) - \['L2CIRCUIT'\]: Resource type.


