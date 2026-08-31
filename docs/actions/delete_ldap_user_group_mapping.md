---
page_title: "gigavuecore_delete_ldap_user_group_mapping Action - gigavuecore"
subcategory: ""
description: |-
  Delete System LDAP User Group Mapping
---

# gigavuecore_delete_ldap_user_group_mapping Action

Delete System LDAP User Group Mapping

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

