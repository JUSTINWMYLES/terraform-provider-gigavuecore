---
page_title: "gigavuecore_get_all_apps_exporter Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Exporter
---

# gigavuecore_get_all_apps_exporter Data Source

Get all Apps Exporter

## Example Usage

```terraform
data "gigavuecore_get_all_apps_exporter" "example" {
  cluster_id = null
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

* `alias` (String) - Alias of the exporter
* `description` (String) - Comments if necessary
* `destination` (Attributes) (see [below for nested schema](#nestedatt--items--destination))
* `gs_group_associated` (List of String)
* `source` (Attributes) (see [below for nested schema](#nestedatt--items--source))
* `ssl_profile` (String) - SSL profile alias
* `status` (String)
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--items--tags))
* `tcp_profile` (String) - TCP profile alias
* `type` (String) - Type of Apps that export
<a id="nestedatt--items--destination"></a>
### Nested Schema for `items.destination`

Read-Only:

* `l3` (Attributes) (see [below for nested schema](#nestedatt--items--destination--l3))
* `l4` (Attributes) (see [below for nested schema](#nestedatt--items--destination--l4))
<a id="nestedatt--items--destination--l3"></a>
### Nested Schema for `items.destination.l3`

Read-Only:

* `ip` (Attributes) (see [below for nested schema](#nestedatt--items--destination--l3--ip))
* `protocol` (String) - Protocol used (when it's auto, it's determined by the App based on context or by discovery)
<a id="nestedatt--items--destination--l3--ip"></a>
### Nested Schema for `items.destination.l3.ip`

Read-Only:

* `dscp` (Number) - DSCP Value to use
* `ttl` (Number) - TTL Value to use
* `ver4` (String) - IPv4 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
* `ver6` (String) - IPv6 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
<a id="nestedatt--items--destination--l4"></a>
### Nested Schema for `items.destination.l4`

Read-Only:

* `port` (Number) - Base port used to export, port is optional for type:gtp-cups
* `protocol` (String) - Protocol used - TCP or UDP
<a id="nestedatt--items--source"></a>
### Nested Schema for `items.source`

Read-Only:

* `interface` (String) - Alias of IP Interface
* `l4_port` (Number) - Base source port number to use for outgoing connections
<a id="nestedatt--items--tags"></a>
### Nested Schema for `items.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

