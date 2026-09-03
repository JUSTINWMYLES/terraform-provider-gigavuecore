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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tunnel_lb_endpoint.example {te_id}/{cluster_id}
```
