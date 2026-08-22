---
page_title: "gigavuecore_proxy_server_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Get Apps Proxy Server Profile for given alias
---

# gigavuecore_proxy_server_profile Resource

Get Apps Proxy Server Profile for given alias

## Example Usage

```terraform
resource "gigavuecore_proxy_server_profile" "example" {
  alias = null
  auth_type = null
  comment = null
  password = null
  periodic_ping = null
  periodic_ping_failure_retry = null
  periodic_ping_interval = null
  periodic_ping_type = null
  port = null
  protocol = null
  proxy_address = null
  ssl_apps = {}
  username = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `auth_type` (String, required)
* `comment` (String, optional)
* `password` (String, optional)
* `periodic_ping` (String, required)
* `periodic_ping_failure_retry` (Number, optional)
* `periodic_ping_interval` (Number, optional)
* `periodic_ping_type` (String, optional)
* `port` (Number, required)
* `protocol` (String, required)
* `proxy_address` (String, required)
* `ssl_apps` (Object({cluster_name}), optional)
  * `cluster_name` (List(String), optional) - Cluster Name where the proxy deployed
* `username` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `id` (String, computed)
* `password` (String, computed)
* `periodic_ping_failure_retry` (Number, computed)
* `periodic_ping_interval` (Number, computed)
* `periodic_ping_type` (String, computed)
* `ssl_apps` (Object({cluster_name}), computed)
  * `cluster_name` (List(String), optional) - Cluster Name where the proxy deployed
* `username` (String, computed)

