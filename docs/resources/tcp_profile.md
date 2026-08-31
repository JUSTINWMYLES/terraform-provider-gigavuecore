---
page_title: "gigavuecore_tcp_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Get Apps TCP Profile for given alias
---

# gigavuecore_tcp_profile Resource

Get Apps TCP Profile for given alias

## Example Usage

```terraform
resource "gigavuecore_tcp_profile" "example" {
  alias            = "example"
  cluster_id       = "example"
  keep_alive_timer = 30
  selective_ack    = "enable"
  syn_retries      = 1
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - Target Cluster ID
* `keep_alive_timer` (Number, optional)
* `selective_ack` (String, optional)
* `syn_retries` (Number, optional)

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
terraform import gigavuecore_tcp_profile.example {alias}:{cluster_id}
```
