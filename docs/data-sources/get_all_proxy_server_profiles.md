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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `apps_proxy_server_profiles` (List(Object({alias, auth_type, comment, password, periodic_ping, periodic_ping_failure_retry, periodic_ping_interval, periodic_ping_type, port, protocol, proxy_address, ssl_apps, username})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

