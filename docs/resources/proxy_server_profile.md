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
  alias                       = "example"
  auth_type                   = "example"
  comment                     = "example"
  password                    = "example"
  periodic_ping               = "example"
  periodic_ping_failure_retry = 0
  periodic_ping_interval      = 0
  periodic_ping_type          = "example"
  port                        = 0
  protocol                    = "example"
  proxy_address               = "example"
  ssl_apps = {
    cluster_name = [ "example" ]
  }
  username = "example"
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
* `ssl_apps` (Attributes, optional) (see [below for nested schema](#nestedatt--ssl_apps))
* `username` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed)
* `password` (String, computed)
* `periodic_ping_failure_retry` (Number, computed)
* `periodic_ping_interval` (Number, computed)
* `periodic_ping_type` (String, computed)
* `ssl_apps` (Attributes, computed) (see [below for nested schema](#nestedatt--ssl_apps))
* `username` (String, computed)

<a id="nestedatt--ssl_apps"></a>
### Nested Schema for `ssl_apps`

Optional:

* `cluster_name` (List of String) - Cluster Name where the proxy deployed

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_proxy_server_profile.example {alias}
```
