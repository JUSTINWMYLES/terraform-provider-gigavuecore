---
page_title: "gigavuecore_v3_user Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new SNMPv3 User
---

# gigavuecore_v3_user Resource

Create a new SNMPv3 User

## Example Usage

```terraform
resource "gigavuecore_v3_user" "example" {
  auth_key      = "example"
  auth_protocol = "md5"
  cluster_id    = "example"
  enabled       = true
  priv_key      = "example"
  priv_protocol = "des"
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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_v3_user.example {username}/{cluster_id}
```
