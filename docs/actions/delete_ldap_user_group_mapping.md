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
    remote_map_table = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `remote_map_table` (List(Dynamic), required)
