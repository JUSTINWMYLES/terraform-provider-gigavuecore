---
page_title: "gigavuecore_inline_ssl_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Load inline SSL profile
---

# gigavuecore_inline_ssl_profile Resource

Load inline SSL profile

## Example Usage

```terraform
resource "gigavuecore_inline_ssl_profile" "example" {
  alias = null
  certificate = {}
  cluster_id = null
  decrypt = {}
  default_action = null
  high_avail = {}
  inbound_tool_early_inspect = {}
  key_map = []
  monitor = null
  network_group = {}
  no_decrypt = {}
  non_ssl_tcp = {}
  one_arm = null
  resilient_inline = {}
  rules = []
  split_proxy = {}
  start_tls = {}
  tcp = {}
  tool = {}
  tool_l3 = {}
  url_cache = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `certificate` (Object({expired, invalid, revocation, self_signed, unknown_ca}), optional)
  * `expired` (String, optional) - SSL profile on expired certificate
  * `invalid` (String, optional) - SSL profile on invalid certificate
  * `revocation` (Object({crl, ocsp}), optional)
    * `crl` (Object({defer, enabled, fail}), optional)
      * `defer` (Number, optional) - timeout in seconds
      * `enabled` (Bool, optional)
      * `fail` (String, optional) - only applicable when enabled
    * `ocsp` (Object({defer, enabled, fail}), optional)
      * `defer` (Number, optional) - timeout in seconds
      * `enabled` (Bool, optional)
      * `fail` (String, optional) - only applicable when enabled
  * `self_signed` (String, optional) - SSL profile on self-signed certificate
  * `unknown_ca` (String, optional) - SSL profile on unknown CA certificate
* `cluster_id` (String, optional) - id of the defining cluster
* `decrypt` (Object({tcp, tool_bypass}), optional)
  * `tcp` (Object({inactive_timeout, port_map}), optional)
    * `inactive_timeout` (Number, optional) - SSL decryption TCP inactive timeout (in minutes)
    * `port_map` (Object({default_out_port, ports}), optional)
      * `default_out_port` (Number, optional) - egress port for decryption port map. 0 is disabled.
      * `ports` (List(Object({in_port, out_port, rule_id})), optional)
  * `tool_bypass` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `default_action` (String, optional) - Action to take if none of the profile rules match
* `high_avail` (Object({active_standby}), optional)
  * `active_standby` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `inbound_tool_early_inspect` (Object({connection_timeout, mode}), optional)
  * `connection_timeout` (Number, optional) - connection timeout timeout in seconds.
  * `mode` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `key_map` (List(Object({hostname, key, rule_id})), optional)
* `monitor` (String, optional)
* `network_group` (Object({multiple_entry}), optional)
  * `multiple_entry` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `no_decrypt` (Object({tool_bypass}), optional)
  * `tool_bypass` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `non_ssl_tcp` (Object({tool_bypass}), optional)
  * `tool_bypass` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `one_arm` (String, optional)
* `resilient_inline` (Object({mode}), optional)
  * `mode` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `rules` (List(Dynamic), optional) - inline SSL profile rules
* `split_proxy` (Object({mode, server_non_pfs_ciphers}), optional)
  * `mode` (Object({enable}), optional)
    * `enable` (Bool, optional)
  * `server_non_pfs_ciphers` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `start_tls` (Object({l4_port}), optional)
  * `l4_port` (List(Number), optional)
* `tcp` (Object({delayed_ack, syn_retries, timewait_timeout}), optional)
  * `delayed_ack` (Bool, optional) - enable/disable TCP delayed ACK
  * `syn_retries` (Number, optional) - TCP Sync retries
  * `timewait_timeout` (Number, optional) - TCP Wait Timeout value
* `tool` (Object({early_engage, fail_action}), optional)
  * `early_engage` (Bool, optional) - enable/disable tool early engage
  * `fail_action` (String, optional) - Action to take if the tool fails
* `tool_l3` (Object({cache_server_cert_timeout, http2_downgrade, nat_pat}), optional)
  * `cache_server_cert_timeout` (Number, optional) - cache server timeout in seconds.
  * `http2_downgrade` (Object({enable}), optional)
    * `enable` (Bool, optional)
  * `nat_pat` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `url_cache` (Object({miss_action, timeout}), optional)
  * `miss_action` (String, optional) - The action to take if local URL category resolution misses
  * `timeout` (Number, optional) - defer timeout in seconds. Only applicable for missAction 'defer'

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `certificate` (Object({expired, invalid, revocation, self_signed, unknown_ca}), computed)
  * `expired` (String, optional) - SSL profile on expired certificate
  * `invalid` (String, optional) - SSL profile on invalid certificate
  * `revocation` (Object({crl, ocsp}), optional)
    * `crl` (Object({defer, enabled, fail}), optional)
      * `defer` (Number, optional) - timeout in seconds
      * `enabled` (Bool, optional)
      * `fail` (String, optional) - only applicable when enabled
    * `ocsp` (Object({defer, enabled, fail}), optional)
      * `defer` (Number, optional) - timeout in seconds
      * `enabled` (Bool, optional)
      * `fail` (String, optional) - only applicable when enabled
  * `self_signed` (String, optional) - SSL profile on self-signed certificate
  * `unknown_ca` (String, optional) - SSL profile on unknown CA certificate
* `cluster_id` (String, computed) - id of the defining cluster
* `decrypt` (Object({tcp, tool_bypass}), computed)
  * `tcp` (Object({inactive_timeout, port_map}), optional)
    * `inactive_timeout` (Number, optional) - SSL decryption TCP inactive timeout (in minutes)
    * `port_map` (Object({default_out_port, ports}), optional)
      * `default_out_port` (Number, optional) - egress port for decryption port map. 0 is disabled.
      * `ports` (List(Object({in_port, out_port, rule_id})), optional)
  * `tool_bypass` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `default_action` (String, computed) - Action to take if none of the profile rules match
* `high_avail` (Object({active_standby}), computed)
  * `active_standby` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `id` (String, computed)
* `inbound_tool_early_inspect` (Object({connection_timeout, mode}), computed)
  * `connection_timeout` (Number, optional) - connection timeout timeout in seconds.
  * `mode` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `key_map` (List(Object({hostname, key, rule_id})), computed)
* `monitor` (String, computed)
* `network_group` (Object({multiple_entry}), computed)
  * `multiple_entry` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `no_decrypt` (Object({tool_bypass}), computed)
  * `tool_bypass` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `non_ssl_tcp` (Object({tool_bypass}), computed)
  * `tool_bypass` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `one_arm` (String, computed)
* `resilient_inline` (Object({mode}), computed)
  * `mode` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `rules` (List(Dynamic), computed) - inline SSL profile rules
* `split_proxy` (Object({mode, server_non_pfs_ciphers}), computed)
  * `mode` (Object({enable}), optional)
    * `enable` (Bool, optional)
  * `server_non_pfs_ciphers` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `start_tls` (Object({l4_port}), computed)
  * `l4_port` (List(Number), optional)
* `tcp` (Object({delayed_ack, syn_retries, timewait_timeout}), computed)
  * `delayed_ack` (Bool, optional) - enable/disable TCP delayed ACK
  * `syn_retries` (Number, optional) - TCP Sync retries
  * `timewait_timeout` (Number, optional) - TCP Wait Timeout value
* `tool` (Object({early_engage, fail_action}), computed)
  * `early_engage` (Bool, optional) - enable/disable tool early engage
  * `fail_action` (String, optional) - Action to take if the tool fails
* `tool_l3` (Object({cache_server_cert_timeout, http2_downgrade, nat_pat}), computed)
  * `cache_server_cert_timeout` (Number, optional) - cache server timeout in seconds.
  * `http2_downgrade` (Object({enable}), optional)
    * `enable` (Bool, optional)
  * `nat_pat` (Object({enable}), optional)
    * `enable` (Bool, optional)
* `url_cache` (Object({miss_action, timeout}), computed)
  * `miss_action` (String, optional) - The action to take if local URL category resolution misses
  * `timeout` (Number, optional) - defer timeout in seconds. Only applicable for missAction 'defer'

