---
page_title: "gigavuecore_redefine_gs_group_ssl_decrypt_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's SSL Decryption Params
---

# gigavuecore_redefine_gs_group_ssl_decrypt_params Action

Redefine GS Group's SSL Decryption Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_ssl_decrypt_params" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    decrypt_fail_action = "example"
    enabled = true
    hsm_pkcs11 = null
    hsm_timeout = 1
    key_cache_timeout = 1
    key_map = "example"
    non_ssl_traffic = "example"
    pending_session_timeout = 1
    session_timeout = 1
    tcp_syn_timeout = 1
    ticket_cache_timeout = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `decrypt_fail_action` (String, optional)
* `enabled` (Bool, optional)
* `hsm_pkcs11` (Dynamic, optional) - GsGroup Ssl Decrypt Hsm Pkcs11 Parameters
* `hsm_timeout` (Number, optional) - in milliseconds
* `key_cache_timeout` (Number, optional) - in seconds
* `key_map` (String, required) - references one of the pre-defined 'SslDecryptionKeyMap' groups
* `non_ssl_traffic` (String, optional)
* `pending_session_timeout` (Number, optional) - in seconds
* `session_timeout` (Number, optional) - in seconds
* `tcp_syn_timeout` (Number, optional) - in seconds
* `ticket_cache_timeout` (Number, optional) - in seconds
