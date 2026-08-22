---
page_title: "gigavuecore_exporter Resource - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_exporter Resource

new in H 5.8

## Example Usage

```terraform
resource "gigavuecore_exporter" "example" {
  alias = null
  description = null
  destination = {}
  gs_group_associated = []
  source = {}
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

* `alias` (String, required) - Alias of the exporter
* `description` (String, optional) - Comments if necessary
* `destination` (Object({l3, l4}), optional)
  * `l3` (Object({ip, protocol}), optional)
    * `ip` (Object({dscp, ttl, ver4, ver6}), optional)
      * `dscp` (Number, optional) - DSCP Value to use
      * `ttl` (Number, optional) - TTL Value to use
      * `ver4` (String, optional) - IPv4 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
      * `ver6` (String, optional) - IPv6 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
    * `protocol` (String, optional) - Protocol used (when it's auto, it's determined by the App based on context or by discovery)
  * `l4` (Object({port, protocol}), optional)
    * `port` (Number, optional) - Base port used to export, port is optional for type:gtp-cups
    * `protocol` (String, required) - Protocol used - TCP or UDP
* `gs_group_associated` (List(String), optional)
* `source` (Object({interface, l4_port}), optional)
  * `interface` (String, required) - Alias of IP Interface
  * `l4_port` (Number, required) - Base source port number to use for outgoing connections
* `ssl_profile` (String, optional) - SSL profile alias
* `status` (String, optional)
* `tags` (List(Object({tag_key, tag_values})), optional)
* `tcp_profile` (String, optional) - TCP profile alias
* `type` (String, optional) - Type of Apps that export

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - Comments if necessary
* `destination` (Object({l3, l4}), computed)
  * `l3` (Object({ip, protocol}), optional)
    * `ip` (Object({dscp, ttl, ver4, ver6}), optional)
      * `dscp` (Number, optional) - DSCP Value to use
      * `ttl` (Number, optional) - TTL Value to use
      * `ver4` (String, optional) - IPv4 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
      * `ver6` (String, optional) - IPv6 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
    * `protocol` (String, optional) - Protocol used (when it's auto, it's determined by the App based on context or by discovery)
  * `l4` (Object({port, protocol}), optional)
    * `port` (Number, optional) - Base port used to export, port is optional for type:gtp-cups
    * `protocol` (String, required) - Protocol used - TCP or UDP
* `gs_group_associated` (List(String), computed)
* `source` (Object({interface, l4_port}), computed)
  * `interface` (String, required) - Alias of IP Interface
  * `l4_port` (Number, required) - Base source port number to use for outgoing connections
* `ssl_profile` (String, computed) - SSL profile alias
* `status` (String, computed)
* `tags` (List(Object({tag_key, tag_values})), computed)
* `tcp_profile` (String, computed) - TCP profile alias
* `type` (String, computed) - Type of Apps that export

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_exporter.example {alias}
```
