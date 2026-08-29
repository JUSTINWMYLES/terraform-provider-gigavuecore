---
page_title: "gigavuecore_tunnel_lb_endpoint Resource - gigavuecore"
subcategory: ""
description: |-
  Load Tunnel Endpoint by id
---

# gigavuecore_tunnel_lb_endpoint Resource

Load Tunnel Endpoint by id

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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `alias` (String, computed) - tunnel endpoint alias
* `ip_address` (String, computed) - tunnel endpoint remote ip address. IPv4 or IPv6


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tunnel_lb_endpoint.example {te_id}
```
