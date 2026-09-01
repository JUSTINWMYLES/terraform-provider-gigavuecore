---
page_title: "gigavuecore_delete_ldap_user_group_mapping Action - gigavuecore"
subcategory: ""
description: |-
  Delete System LDAP User Group Mapping
---

# gigavuecore_delete_ldap_user_group_mapping Action

Delete System LDAP User Group Mapping

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_ldap_user_group_mapping" "example" {
  config {
    cluster_id = "example"
    remote_map_table = [{
      local_account_name = "example"
      remote_base_dn     = "example"
    }]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `remote_map_table` (Attributes List, required) (see [below for nested schema](#nestedatt--remote_map_table))

<a id="nestedatt--remote_map_table"></a>
### Nested Schema for `remote_map_table`

Optional:

* `local_account_name` (String) - Specifies local account to which remote base-dn needs to be mapped.
* `remote_base_dn` (String) - Specifies the base-dn of the remote user group to be mapped to local account

