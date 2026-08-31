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
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  destination = {
    l3 = {
      ip = {
        dscp = 0
        ttl  = 1
        ver4 = "example"
        ver6 = "example"
      }
      protocol = "ipv4"
    }
    l4 = {
      port     = 1
      protocol = "tcp"
    }
  }
  gs_group_associated = [ "example" ]
  source = {
    interface = "example"
    l4_port   = 1
  }
  ssl_profile = "example"
  status      = "active"
  tags = [{
    tag_key    = "example"
    tag_values = [ "example" ]
  }]
  tcp_profile = "example"
  type        = "mobility-cups"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the exporter
* `cluster_id` (String, required) - Target Cluster ID
* `description` (String, optional) - Comments if necessary
* `destination` (Attributes, optional) (see [below for nested schema](#nestedatt--destination))
* `gs_group_associated` (List of String, optional)
* `source` (Attributes, optional) (see [below for nested schema](#nestedatt--source))
* `ssl_profile` (String, optional) - SSL profile alias
* `status` (String, optional)
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))
* `tcp_profile` (String, optional) - TCP profile alias
* `type` (String, optional) - Type of Apps that export

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_exporter.example {alias}/{cluster_id}
```
