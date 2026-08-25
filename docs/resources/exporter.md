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
  alias               = null
  description         = null
  destination         = {}
  gs_group_associated = []
  source              = {}
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

* `alias` (String, required) - Alias of the exporter
* `description` (String, optional) - Comments if necessary
* `destination` (Attributes, optional) (see [below for nested schema](#nestedatt--destination))
* `gs_group_associated` (List of String, optional)
* `source` (Attributes, optional) (see [below for nested schema](#nestedatt--source))
* `ssl_profile` (String, optional) - SSL profile alias
* `status` (String, optional)
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))
* `tcp_profile` (String, optional) - TCP profile alias
* `type` (String, optional) - Type of Apps that export

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed) - Comments if necessary
* `destination` (Attributes, computed) (see [below for nested schema](#nestedatt--destination))
* `gs_group_associated` (List of String, computed)
* `source` (Attributes, computed) (see [below for nested schema](#nestedatt--source))
* `ssl_profile` (String, computed) - SSL profile alias
* `status` (String, computed)
* `tags` (Attributes List, computed) (see [below for nested schema](#nestedatt--tags))
* `tcp_profile` (String, computed) - TCP profile alias
* `type` (String, computed) - Type of Apps that export

<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Optional:

* `l3` (Attributes) (see [below for nested schema](#nestedatt--destination--l3))
* `l4` (Attributes) (see [below for nested schema](#nestedatt--destination--l4))
<a id="nestedatt--destination--l3"></a>
### Nested Schema for `destination.l3`

Optional:

* `ip` (Attributes) (see [below for nested schema](#nestedatt--destination--l3--ip))
* `protocol` (String) - Protocol used (when it's auto, it's determined by the App based on context or by discovery)
<a id="nestedatt--destination--l3--ip"></a>
### Nested Schema for `destination.l3.ip`

Optional:

* `dscp` (Number) - DSCP Value to use
* `ttl` (Number) - TTL Value to use
* `ver4` (String) - IPv4 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
* `ver6` (String) - IPv6 address. If not specified, it will be determined by the App based on other traffic correlation or external means (auto mode).
<a id="nestedatt--destination--l4"></a>
### Nested Schema for `destination.l4`

Required:

* `protocol` (String) - Protocol used - TCP or UDP
Optional:

* `port` (Number) - Base port used to export, port is optional for type:gtp-cups
<a id="nestedatt--source"></a>
### Nested Schema for `source`

Required:

* `interface` (String) - Alias of IP Interface
* `l4_port` (Number) - Base source port number to use for outgoing connections
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_exporter.example {alias}
```
