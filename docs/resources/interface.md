---
page_title: "gigavuecore_interface Resource - gigavuecore"
subcategory: ""
description: |-
  Configure IP Interfaces
---

# gigavuecore_interface Resource

Configure IP Interfaces

## Example Usage

```terraform
resource "gigavuecore_interface" "example" {
  alias             = "example"
  attach            = ["example"]
  cluster_id        = "example"
  comment           = "example"
  gateway           = "example"
  gs_groups         = ["example"]
  hw_address        = "example"
  ip_address        = "example"
  ip_mask           = "example"
  mtu               = 1500
  netflow_exporters = ["example"]
  tags = [{
    tag_key    = "example"
    tag_values = ["example"]
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - ip interface name
* `attach` (List of String, optional) - network ports ,tool ports or circuit ports
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `gateway` (String, optional) - gateway ipv4 or ipv6 address
* `gs_groups` (List of String, optional) - Gs Groups associated with the IP Interface
* `hw_address` (String, optional)
* `ip_address` (String, optional) - ipv4/ipv6 address
* `ip_mask` (String, optional) - ipAddress netmask required with ipAddress
* `mtu` (Number, optional)
* `netflow_exporters` (List of String, optional) - Netflow Exporters associated with the IP Interface
* `tags` (Attributes List, optional) (see [below for nested schema](#nestedatt--tags))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `ip_type` (String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
terraform import gigavuecore_interface.example {alias}/{cluster_id}
```
