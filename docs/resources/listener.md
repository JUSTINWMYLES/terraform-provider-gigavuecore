---
page_title: "gigavuecore_listener Resource - gigavuecore"
subcategory: ""
description: |-
  Create Apps Listener
---

# gigavuecore_listener Resource

Create Apps Listener

## Example Usage

```terraform
resource "gigavuecore_listener" "example" {
  alias               = "example"
  cluster_id          = "example"
  description         = "example"
  gs_group_associated = ["example"]
  ip_interface        = ["example"]
  l3 = {
    dscp     = 0
    protocol = "ipv4"
    ttl      = 1
  }
  l4 = {
    port_list = [1]
    protocol  = "tcp"
  }
  mode        = "promiscuous"
  ssl_profile = "example"
  status      = "active"
  tags = [{
    tag_key    = "example"
    tag_values = ["example"]
  }]
  tcp_profile = "example"
  type        = "gtp-cups"
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
terraform import gigavuecore_listener.example {alias}/{cluster_id}
```
