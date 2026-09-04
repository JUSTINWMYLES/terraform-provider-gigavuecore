---
page_title: "gigavuecore_add_fm_template_child Action - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.13.01
---

# gigavuecore_add_fm_template_child Action

new in FM 5.13.01

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_add_fm_template_child" "example" {
  config {
    body_child_type  = "ldapServers"
    body_config_type = "LDAP_SERVERS_TEMPLATE"
    child_type       = "example"
    config = {
      ldap_servers = [{
        order          = "example"
        server_address = "example"
      }]
      remote_map_table = [{
        local_account_name = "example"
        remote_base_dn     = "example"
      }]
    }
    config_level       = "GLOBAL"
    config_level_value = ["example"]
    config_type        = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `body_child_type` (String, optional) - Child Type of the FM template Configuration
* `body_config_type` (String, optional) - Configuration Type of the FM template
* `child_type` (String, required) - childType of the fm template
* `config` (Attributes, optional) (see [below for nested schema](#nestedatt--config))
* `config_level` (String, optional) - Scope of the applied FM template
* `config_level_value` (List of String, optional)
* `config_type` (String, required) - configType of the fm template

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

* `ldap_servers` (Attributes List) - List of LDAP Servers. Available for ConfigType LDAP\_SERVERS\_TEMPLATE ChildType ldapServers (see [below for nested schema](#nestedatt--config--ldap_servers))
* `remote_map_table` (Attributes List) - List of LDAP User Group Mapping. Available for ConfigType LDAP\_SYSTEM\_CONFIG\_TEMPLATE ChildType remoteMapTable (see [below for nested schema](#nestedatt--config--remote_map_table))

<a id="nestedatt--config--ldap_servers"></a>
### Nested Schema for `config.ldap_servers`

Required:

* `server_address` (String) - ipv4 or ipv6 or hostname. Specifies address of the LDAP server where authentication requests will be sent

Optional:

* `order` (String) - The order in which the server is to be reached. 1 means server will be contacted first

<a id="nestedatt--config--remote_map_table"></a>
### Nested Schema for `config.remote_map_table`

Optional:

* `local_account_name` (String) - Specifies local account to which remote base-dn needs to be mapped.
* `remote_base_dn` (String) - Specifies the base-dn of the remote user group to be mapped to local account

