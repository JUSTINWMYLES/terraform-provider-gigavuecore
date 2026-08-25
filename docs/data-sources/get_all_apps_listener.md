---
page_title: "gigavuecore_get_all_apps_listener Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Listener
---

# gigavuecore_get_all_apps_listener Data Source

Get all Apps Listener

## Example Usage

```terraform
data "gigavuecore_get_all_apps_listener" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `apps_listeners` (Attributes List, computed) (see [below for nested schema](#nestedatt--apps_listeners))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--apps_listeners"></a>
### Nested Schema for `apps_listeners`

Read-Only:

* `alias` (String) - Alias of the listener
* `description` (String) - Comments if necessary
* `gs_group_associated` (List of String)
* `ip_interface` (List of String)
* `l3` (Attributes) (see [below for nested schema](#nestedatt--apps_listeners--l3))
* `l4` (Attributes) (see [below for nested schema](#nestedatt--apps_listeners--l4))
* `mode` (String) - Listen to IP interface or promiscuous mode
* `ssl_profile` (String) - SSL profile
* `status` (String)
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--apps_listeners--tags))
* `tcp_profile` (String) - TCP profile
* `type` (String) - Type of Apps that listen
<a id="nestedatt--apps_listeners--l3"></a>
### Nested Schema for `apps_listeners.l3`

Read-Only:

* `dscp` (Number) - DSCP Value to use
* `protocol` (String) - Protocol used - ipv4 , ipv6 or both
* `ttl` (Number) - TTL Value to use
<a id="nestedatt--apps_listeners--l4"></a>
### Nested Schema for `apps_listeners.l4`

Read-Only:

* `port_list` (List of Number) - Port lists to listen on
* `protocol` (String) - Protocol used - TCP or UDP
<a id="nestedatt--apps_listeners--tags"></a>
### Nested Schema for `apps_listeners.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

