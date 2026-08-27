---
page_title: "gigavuecore_get_all_proxy_server_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Proxy Server Profile
---

# gigavuecore_get_all_proxy_server_profiles Data Source

Get all Apps Proxy Server Profile

## Example Usage

```terraform
data "gigavuecore_get_all_proxy_server_profiles" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String)
* `auth_type` (String)
* `comment` (String)
* `password` (String)
* `periodic_ping` (String)
* `periodic_ping_failure_retry` (Number)
* `periodic_ping_interval` (Number)
* `periodic_ping_type` (String)
* `port` (Number)
* `protocol` (String)
* `proxy_address` (String)
* `ssl_apps` (Attributes) (see [below for nested schema](#nestedatt--items--ssl_apps))
* `username` (String)
<a id="nestedatt--items--ssl_apps"></a>
### Nested Schema for `items.ssl_apps`

Read-Only:

* `cluster_name` (List of String) - Cluster Name where the proxy deployed

