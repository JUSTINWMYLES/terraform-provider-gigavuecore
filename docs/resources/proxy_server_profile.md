---
page_title: "gigavuecore_proxy_server_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Create Apps Proxy Server Profile
---

# gigavuecore_proxy_server_profile Resource

Create Apps Proxy Server Profile

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
    cluster_name = ["example"]
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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_proxy_server_profile.example {alias}
```
