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

* `apps_proxy_server_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--apps_proxy_server_profiles))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--apps_proxy_server_profiles"></a>
### Nested Schema for `apps_proxy_server_profiles`

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
* `ssl_apps` (Attributes) (see [below for nested schema](#nestedatt--apps_proxy_server_profiles--ssl_apps))
* `username` (String)
<a id="nestedatt--apps_proxy_server_profiles--ssl_apps"></a>
### Nested Schema for `apps_proxy_server_profiles.ssl_apps`

Read-Only:

* `cluster_name` (List of String) - Cluster Name where the proxy deployed
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

