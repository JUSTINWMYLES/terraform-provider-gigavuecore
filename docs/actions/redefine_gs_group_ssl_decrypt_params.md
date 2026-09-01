---
page_title: "gigavuecore_redefine_gs_group_ssl_decrypt_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's SSL Decryption Params
---

# gigavuecore_redefine_gs_group_ssl_decrypt_params Action

Redefine GS Group's SSL Decryption Params

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_ssl_decrypt_params" "example" {
  config {
    alias               = "example"
    cluster_id          = "example"
    decrypt_fail_action = "drop"
    enabled             = true
    hsm_pkcs11 = {
      debug_level    = 0
      dynamic_object = true
      load_sharing   = true
    }
    hsm_timeout             = 2
    key_cache_timeout       = 1
    key_map                 = "example"
    non_ssl_traffic         = "drop"
    pending_session_timeout = 30
    session_timeout         = 30
    tcp_syn_timeout         = 20
    ticket_cache_timeout    = 1
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `decrypt_fail_action` (String, optional)
* `enabled` (Boolean, optional)
* `hsm_pkcs11` (Attributes, optional) - GsGroup Ssl Decrypt Hsm Pkcs11 Parameters (see [below for nested schema](#nestedatt--hsm_pkcs11))
* `hsm_timeout` (Number, optional) - in milliseconds
* `key_cache_timeout` (Number, optional) - in seconds
* `key_map` (String, required) - references one of the pre-defined 'SslDecryptionKeyMap' groups
* `non_ssl_traffic` (String, optional)
* `pending_session_timeout` (Number, optional) - in seconds
* `session_timeout` (Number, optional) - in seconds
* `tcp_syn_timeout` (Number, optional) - in seconds
* `ticket_cache_timeout` (Number, optional) - in seconds

<a id="nestedatt--hsm_pkcs11"></a>
### Nested Schema for `hsm_pkcs11`

Optional:

* `debug_level` (Number) - hsm pkcs11 debug level
* `dynamic_object` (Boolean) - hsm pkcs11 dynamic object
* `load_sharing` (Boolean) - hsm pkcs11 load sharing

