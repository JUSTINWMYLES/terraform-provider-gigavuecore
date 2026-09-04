---
page_title: "gigavuecore_tunnel_lb_endpoint Resource - gigavuecore"
subcategory: ""
description: |-
  Create a Tunnel Endpoint
---

# gigavuecore_tunnel_lb_endpoint Resource

Create a Tunnel Endpoint

## Example Usage

```terraform
resource "gigavuecore_tunnel_lb_endpoint" "example" {
  alias      = "example"
  cluster_id = "example"
  ip_address = "example"
  te_id      = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - tunnel endpoint alias
* `cluster_id` (String, required) - Target Cluster ID
* `ip_address` (String, optional) - tunnel endpoint remote ip address. IPv4 or IPv6
* `te_id` (String, required) - tunnel endpoint alias, format: teN , where 1 <= N <= 128

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
terraform import gigavuecore_tunnel_lb_endpoint.example {te_id}/{cluster_id}
```
