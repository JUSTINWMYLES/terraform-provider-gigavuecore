---
page_title: "gigavuecore_tcp_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Create Apps TCP Profile
---

# gigavuecore_tcp_profile Resource

Create Apps TCP Profile

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tcp_profile.example {alias}:{cluster_id}
```
