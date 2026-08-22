---
page_title: "gigavuecore_get_system_web_configuration Data Source - gigavuecore"
subcategory: ""
description: |-
  get system web configuration
---

# gigavuecore_get_system_web_configuration Data Source

get system web configuration

## Example Usage

```terraform
data "gigavuecore_get_system_web_configuration" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - Target Cluster Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auto_logout_timeout` (Number, computed) - Auto Logout Timeout in seconds
* `cluster_name` (String, computed) - cluster name
* `enable_https` (Bool, computed) - enableHttps
* `http_port` (Number, computed) - Http Port Number
* `https_port` (Number, computed) - Https Port Number
* `server_ssl_min_version` (String, computed) - SSL/TLS minimum version supported
* `session_renewal` (Number, computed) - Session renewal in seconds
* `session_timeout` (Number, computed) - Session renewal in seconds

