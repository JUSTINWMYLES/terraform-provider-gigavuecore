---
page_title: "gigavuecore_v3_user Resource - gigavuecore"
subcategory: ""
description: |-
  Find SNMPv3 User by name
---

# gigavuecore_v3_user Resource

Find SNMPv3 User by name

## Example Usage

```terraform
resource "gigavuecore_v3_user" "example" {
  auth_key      = "example"
  auth_protocol = "example"
  cluster_id    = "example"
  enabled       = true
  priv_key      = "example"
  priv_protocol = "example"
  read_only     = true
  username      = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `auth_key` (String, optional)
* `auth_protocol` (String, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional)
* `priv_key` (String, optional)
* `priv_protocol` (String, optional)
* `read_only` (Boolean, optional) - This property is introduced for GUI's purpose to restrict the edit/delete on the SNMPv3 user used by FM. If this property doesn't exist, consider it false. And this property is ignored if it exists in the create/update spec.
* `username` (String, required)

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
terraform import gigavuecore_v3_user.example {username}/{cluster_id}
```
