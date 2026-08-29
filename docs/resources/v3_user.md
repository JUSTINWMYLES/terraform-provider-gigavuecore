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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `auth_key` (String, computed)
* `auth_protocol` (String, computed)
* `enabled` (Boolean, computed)
* `priv_key` (String, computed)
* `priv_protocol` (String, computed)
* `read_only` (Boolean, computed) - This property is introduced for GUI's purpose to restrict the edit/delete on the SNMPv3 user used by FM. If this property doesn't exist, consider it false. And this property is ignored if it exists in the create/update spec.
* `user_name` (String, computed) - name of the target SNMPv3 User
* `username` (String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_v3_user.example {user_name}
```
