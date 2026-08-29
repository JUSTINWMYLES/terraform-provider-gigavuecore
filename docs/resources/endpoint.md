---
page_title: "gigavuecore_endpoint Resource - gigavuecore"
subcategory: ""
description: |-
  Find a configured SSL Decryption Endpoint by alias
---

# gigavuecore_endpoint Resource

Find a configured SSL Decryption Endpoint by alias

## Example Usage

```terraform
resource "gigavuecore_endpoint" "example" {
  address    = "example"
  alias      = "example"
  cluster_id = "example"
  port       = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `address` (String, required) - IPv4 address of the SSL endpoint (server)
* `alias` (String, required) - Alias of the Endpoint
* `cluster_id` (String, required) - Target Cluster ID
* `port` (Number, required) - Optional port of the SSL endpoint (server). Value of 0 indicates any port

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
terraform import gigavuecore_endpoint.example {alias}/{cluster_id}
```
