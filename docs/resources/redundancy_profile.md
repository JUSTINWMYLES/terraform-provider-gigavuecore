---
page_title: "gigavuecore_redundancy_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Find Redundancy Profile by alias
---

# gigavuecore_redundancy_profile Resource

Find Redundancy Profile by alias

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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_redundancy_profile.example {alias}/{cluster_id}
```
