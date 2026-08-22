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
  alias = null
  description = null
  gs_group_associated = []
  ip_interface = []
  l3 = {}
  l4 = {}
  mode = null
  ssl_profile = null
  status = null
  tags = []
  tcp_profile = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the listener
* `description` (String, optional) - Comments if necessary
* `gs_group_associated` (List(String), optional)
* `ip_interface` (List(String), optional)
* `l3` (Object({dscp, protocol, ttl}), optional)
  * `dscp` (Number, optional) - DSCP Value to use
  * `protocol` (String, optional) - Protocol used - ipv4 , ipv6 or both
  * `ttl` (Number, optional) - TTL Value to use
* `l4` (Object({port_list, protocol}), optional)
  * `port_list` (List(Number), required) - Port lists to listen on
  * `protocol` (String, required) - Protocol used - TCP or UDP
* `mode` (String, optional) - Listen to IP interface or promiscuous mode
* `ssl_profile` (String, optional) - SSL profile
* `status` (String, optional)
* `tags` (List(Object({tag_key, tag_values})), optional)
* `tcp_profile` (String, optional) - TCP profile
* `type` (String, optional) - Type of Apps that listen

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - Comments if necessary
* `gs_group_associated` (List(String), computed)
* `ip_interface` (List(String), computed)
* `l3` (Object({dscp, protocol, ttl}), computed)
  * `dscp` (Number, optional) - DSCP Value to use
  * `protocol` (String, optional) - Protocol used - ipv4 , ipv6 or both
  * `ttl` (Number, optional) - TTL Value to use
* `l4` (Object({port_list, protocol}), computed)
  * `port_list` (List(Number), required) - Port lists to listen on
  * `protocol` (String, required) - Protocol used - TCP or UDP
* `mode` (String, computed) - Listen to IP interface or promiscuous mode
* `ssl_profile` (String, computed) - SSL profile
* `status` (String, computed)
* `tags` (List(Object({tag_key, tag_values})), computed)
* `tcp_profile` (String, computed) - TCP profile
* `type` (String, computed) - Type of Apps that listen

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_listener.example {alias}
```
