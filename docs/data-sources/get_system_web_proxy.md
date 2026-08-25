---
page_title: "gigavuecore_get_system_web_proxy Data Source - gigavuecore"
subcategory: ""
description: |-
  get system web proxy
---

# gigavuecore_get_system_web_proxy Data Source

get system web proxy

## Example Usage

```terraform
data "gigavuecore_get_system_web_proxy" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auth_type` (String, computed) - Auth type
* `cluster_name` (String, computed) - cluster name
* `password` (String, computed) - Proxy Password
* `proxy_address` (String, computed) - Proxy address
* `proxy_port` (Number, computed) - Proxy port
* `username` (String, computed) - Proxy Username


