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
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias of the listener
* `description` (String) - Comments if necessary
* `gs_group_associated` (List of String)
* `ip_interface` (List of String)
* `l3` (Attributes) (see [below for nested schema](#nestedatt--items--l3))
* `l4` (Attributes) (see [below for nested schema](#nestedatt--items--l4))
* `mode` (String) - Listen to IP interface or promiscuous mode
* `ssl_profile` (String) - SSL profile
* `status` (String)
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--items--tags))
* `tcp_profile` (String) - TCP profile
* `type` (String) - Type of Apps that listen

<a id="nestedatt--items--l3"></a>
### Nested Schema for `items.l3`

Read-Only:

* `dscp` (Number) - DSCP Value to use
* `protocol` (String) - Protocol used - ipv4 , ipv6 or both
* `ttl` (Number) - TTL Value to use

<a id="nestedatt--items--l4"></a>
### Nested Schema for `items.l4`

Read-Only:

* `port_list` (List of Number) - Port lists to listen on
* `protocol` (String) - Protocol used - TCP or UDP

<a id="nestedatt--items--tags"></a>
### Nested Schema for `items.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

