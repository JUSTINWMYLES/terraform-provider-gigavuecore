---
page_title: "gigavuecore_update_fabric_resource_config Action - gigavuecore"
subcategory: ""
description: |-
  Update fabric resource allocation policy configuration
---

# gigavuecore_update_fabric_resource_config Action

Update fabric resource allocation policy configuration

## Example Usage

```terraform
action "gigavuecore_update_fabric_resource_config" "example" {
  config {
    mode = "example"
    scope = "example"
    type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `mode` (String, optional) - \['SHARE' or 'NOT\_SHARE'\]: Resource sharing mode. SHARE: same resource can be shared by different fabric maps; NOT\_SHARE: not shared. Default: NOT\_SHARE.
* `scope` (String, optional) - \['GLOBAL'\]: Scope of resource pool. GLOBAL: only one global resource pool.
* `type` (String, optional) - \['L2CIRCUIT'\]: Resource type.
