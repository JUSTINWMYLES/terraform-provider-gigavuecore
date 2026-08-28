---
page_title: "gigavuecore_listener Resource - gigavuecore"
subcategory: ""
description: |-
  Get Apps Listener for given alias for GS as Service
---

# gigavuecore_listener Resource

Get Apps Listener for given alias for GS as Service

## Example Usage

```terraform
resource "gigavuecore_listener" "example" {
  alias               = null
  cluster_id          = null
  description         = null
  gs_group_associated = []
  ip_interface        = []
  l3                  = {}
  l4                  = {}
  mode                = null
  ssl_profile         = null
  status              = null
  tags                = []
  tcp_profile         = null
  type                = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the listener
* `cluster_id` (String, required) - Target Cluster ID
* `description` (String, optional) - Comments if necessary
* `gs_group_associated` (List of String, optional)
* `ip_interface` (List of String, optional)
* `l3` (Attributes, optional) (see [below for nested schema](#nestedatt--l3))
* `l4` (Attributes, optional) (see [below for nested schema](#nestedatt--l4))
* `mode` (String, optional) - Listen to IP interface or promiscuous mode
* `ssl_profile` (String, optional) - SSL profile
* `status` (String, optional)
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))
* `tcp_profile` (String, optional) - TCP profile
* `type` (String, optional) - Type of Apps that listen

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - Comments if necessary
* `gs_group_associated` (List of String, computed)
* `ip_interface` (List of String, computed)
* `l3` (Attributes, computed) (see [below for nested schema](#nestedatt--l3))
* `l4` (Attributes, computed) (see [below for nested schema](#nestedatt--l4))
* `mode` (String, computed) - Listen to IP interface or promiscuous mode
* `ssl_profile` (String, computed) - SSL profile
* `status` (String, computed)
* `tags` (Attributes List, computed) (see [below for nested schema](#nestedatt--tags))
* `tcp_profile` (String, computed) - TCP profile
* `type` (String, computed) - Type of Apps that listen

<a id="nestedatt--l3"></a>
### Nested Schema for `l3`

Optional:

* `dscp` (Number) - DSCP Value to use
* `protocol` (String) - Protocol used - ipv4 , ipv6 or both
* `ttl` (Number) - TTL Value to use
<a id="nestedatt--l4"></a>
### Nested Schema for `l4`

Required:

* `port_list` (List of Number) - Port lists to listen on
* `protocol` (String) - Protocol used - TCP or UDP
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_listener.example {alias}
```
