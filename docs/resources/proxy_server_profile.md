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
  auth_type                   = "none"
  comment                     = "example"
  password                    = "example"
  periodic_ping               = "enable"
  periodic_ping_failure_retry = 1
  periodic_ping_interval      = 1
  periodic_ping_type          = "http-connect"
  port                        = 1
  protocol                    = "http"
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--ssl_apps"></a>
### Nested Schema for `ssl_apps`

Optional:

* `cluster_name` (List of String) - Cluster Name where the proxy deployed
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_proxy_server_profile.example {alias}
```
