---
page_title: "gigavuecore_redundancy_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Redundancy Profile
---

# gigavuecore_redundancy_profile Resource

Create a new Redundancy Profile

## Example Usage

```terraform
resource "gigavuecore_redundancy_profile" "example" {
  alias           = "example"
  cluster_id      = "example"
  protection_role = "suspended"
  signaling_port  = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline Network alias. Unique within a cluster
* `cluster_id` (String, required) - Target Cluster ID
* `protection_role` (String, optional) - only applicable for 'protected' inline networks
* `signaling_port` (String, required) - only applicable for 'protected' inline networks

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
terraform import gigavuecore_redundancy_profile.example {alias}/{cluster_id}
```
