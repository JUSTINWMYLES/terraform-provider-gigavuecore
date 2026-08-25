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
  alias                      = null
  certificate                = {}
  cluster_id                 = null
  decrypt                    = {}
  default_action             = null
  high_avail                 = {}
  inbound_tool_early_inspect = {}
  key_map                    = []
  monitor                    = null
  network_group              = {}
  no_decrypt                 = {}
  non_ssl_tcp                = {}
  one_arm                    = null
  resilient_inline           = {}
  rules                      = []
  split_proxy                = {}
  start_tls                  = {}
  tcp                        = {}
  tool                       = {}
  tool_l3                    = {}
  url_cache                  = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `certificate` (Attributes, optional) (see [below for nested schema](#nestedatt--certificate))
* `cluster_id` (String, optional) - id of the defining cluster
* `decrypt` (Attributes, optional) (see [below for nested schema](#nestedatt--decrypt))
* `default_action` (String, optional) - Action to take if none of the profile rules match
* `high_avail` (Attributes, optional) (see [below for nested schema](#nestedatt--high_avail))
* `inbound_tool_early_inspect` (Attributes, optional) (see [below for nested schema](#nestedatt--inbound_tool_early_inspect))
* `key_map` (Attributes List, optional) (see [below for nested schema](#nestedatt--key_map))
* `monitor` (String, optional)
* `network_group` (Attributes, optional) (see [below for nested schema](#nestedatt--network_group))
* `no_decrypt` (Attributes, optional) (see [below for nested schema](#nestedatt--no_decrypt))
* `non_ssl_tcp` (Attributes, optional) (see [below for nested schema](#nestedatt--non_ssl_tcp))
* `one_arm` (String, optional)
* `resilient_inline` (Attributes, optional) (see [below for nested schema](#nestedatt--resilient_inline))
* `rules` (List of Dynamic, optional) - inline SSL profile rules
* `split_proxy` (Attributes, optional) (see [below for nested schema](#nestedatt--split_proxy))
* `start_tls` (Attributes, optional) (see [below for nested schema](#nestedatt--start_tls))
* `tcp` (Attributes, optional) (see [below for nested schema](#nestedatt--tcp))
* `tool` (Attributes, optional) (see [below for nested schema](#nestedatt--tool))
* `tool_l3` (Attributes, optional) (see [below for nested schema](#nestedatt--tool_l3))
* `url_cache` (Attributes, optional) (see [below for nested schema](#nestedatt--url_cache))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `certificate` (Attributes, computed) (see [below for nested schema](#nestedatt--certificate))
* `cluster_id` (String, computed) - id of the defining cluster
* `decrypt` (Attributes, computed) (see [below for nested schema](#nestedatt--decrypt))
* `default_action` (String, computed) - Action to take if none of the profile rules match
* `high_avail` (Attributes, computed) (see [below for nested schema](#nestedatt--high_avail))
* `id` (String, computed)
* `inbound_tool_early_inspect` (Attributes, computed) (see [below for nested schema](#nestedatt--inbound_tool_early_inspect))
* `key_map` (Attributes List, computed) (see [below for nested schema](#nestedatt--key_map))
* `monitor` (String, computed)
* `network_group` (Attributes, computed) (see [below for nested schema](#nestedatt--network_group))
* `no_decrypt` (Attributes, computed) (see [below for nested schema](#nestedatt--no_decrypt))
* `non_ssl_tcp` (Attributes, computed) (see [below for nested schema](#nestedatt--non_ssl_tcp))
* `one_arm` (String, computed)
* `resilient_inline` (Attributes, computed) (see [below for nested schema](#nestedatt--resilient_inline))
* `rules` (List of Dynamic, computed) - inline SSL profile rules
* `split_proxy` (Attributes, computed) (see [below for nested schema](#nestedatt--split_proxy))
* `start_tls` (Attributes, computed) (see [below for nested schema](#nestedatt--start_tls))
* `tcp` (Attributes, computed) (see [below for nested schema](#nestedatt--tcp))
* `tool` (Attributes, computed) (see [below for nested schema](#nestedatt--tool))
* `tool_l3` (Attributes, computed) (see [below for nested schema](#nestedatt--tool_l3))
* `url_cache` (Attributes, computed) (see [below for nested schema](#nestedatt--url_cache))

<a id="nestedatt--certificate"></a>
### Nested Schema for `certificate`

Optional:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) (see [below for nested schema](#nestedatt--certificate--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate
<a id="nestedatt--certificate--revocation"></a>
### Nested Schema for `certificate.revocation`

Optional:

* `crl` (Attributes) (see [below for nested schema](#nestedatt--certificate--revocation--crl))
* `ocsp` (Attributes) (see [below for nested schema](#nestedatt--certificate--revocation--ocsp))
<a id="nestedatt--certificate--revocation--crl"></a>
### Nested Schema for `certificate.revocation.crl`

Optional:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--certificate--revocation--ocsp"></a>
### Nested Schema for `certificate.revocation.ocsp`

Optional:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--decrypt"></a>
### Nested Schema for `decrypt`

Optional:

* `tcp` (Attributes) (see [below for nested schema](#nestedatt--decrypt--tcp))
* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--decrypt--tool_bypass))
<a id="nestedatt--decrypt--tcp"></a>
### Nested Schema for `decrypt.tcp`

Optional:

* `inactive_timeout` (Number) - SSL decryption TCP inactive timeout (in minutes)
* `port_map` (Attributes) (see [below for nested schema](#nestedatt--decrypt--tcp--port_map))
<a id="nestedatt--decrypt--tcp--port_map"></a>
### Nested Schema for `decrypt.tcp.port_map`

Optional:

* `default_out_port` (Number) - egress port for decryption port map. 0 is disabled.
* `ports` (Attributes List) (see [below for nested schema](#nestedatt--decrypt--tcp--port_map--ports))
<a id="nestedatt--decrypt--tcp--port_map--ports"></a>
### Nested Schema for `decrypt.tcp.port_map.ports`

Required:

* `in_port` (Number) - ingress port for decryption port map
* `out_port` (Number) - egress port for decryption port map
Optional:

* `rule_id` (Number)
<a id="nestedatt--decrypt--tool_bypass"></a>
### Nested Schema for `decrypt.tool_bypass`

Optional:

* `enable` (Boolean)
<a id="nestedatt--high_avail"></a>
### Nested Schema for `high_avail`

Optional:

* `active_standby` (Attributes) (see [below for nested schema](#nestedatt--high_avail--active_standby))
<a id="nestedatt--high_avail--active_standby"></a>
### Nested Schema for `high_avail.active_standby`

Optional:

* `enable` (Boolean)
<a id="nestedatt--inbound_tool_early_inspect"></a>
### Nested Schema for `inbound_tool_early_inspect`

Optional:

* `connection_timeout` (Number) - connection timeout timeout in seconds.
* `mode` (Attributes) (see [below for nested schema](#nestedatt--inbound_tool_early_inspect--mode))
<a id="nestedatt--inbound_tool_early_inspect--mode"></a>
### Nested Schema for `inbound_tool_early_inspect.mode`

Optional:

* `enable` (Boolean)
<a id="nestedatt--key_map"></a>
### Nested Schema for `key_map`

Required:

* `hostname` (String) - hostname or IP address
* `key` (String) - SSL key alias
Optional:

* `rule_id` (Number)
<a id="nestedatt--network_group"></a>
### Nested Schema for `network_group`

Optional:

* `multiple_entry` (Attributes) (see [below for nested schema](#nestedatt--network_group--multiple_entry))
<a id="nestedatt--network_group--multiple_entry"></a>
### Nested Schema for `network_group.multiple_entry`

Optional:

* `enable` (Boolean)
<a id="nestedatt--no_decrypt"></a>
### Nested Schema for `no_decrypt`

Optional:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--no_decrypt--tool_bypass))
<a id="nestedatt--no_decrypt--tool_bypass"></a>
### Nested Schema for `no_decrypt.tool_bypass`

Optional:

* `enable` (Boolean)
<a id="nestedatt--non_ssl_tcp"></a>
### Nested Schema for `non_ssl_tcp`

Optional:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--non_ssl_tcp--tool_bypass))
<a id="nestedatt--non_ssl_tcp--tool_bypass"></a>
### Nested Schema for `non_ssl_tcp.tool_bypass`

Optional:

* `enable` (Boolean)
<a id="nestedatt--resilient_inline"></a>
### Nested Schema for `resilient_inline`

Optional:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--resilient_inline--mode))
<a id="nestedatt--resilient_inline--mode"></a>
### Nested Schema for `resilient_inline.mode`

Optional:

* `enable` (Boolean)
<a id="nestedatt--split_proxy"></a>
### Nested Schema for `split_proxy`

Optional:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--split_proxy--mode))
* `server_non_pfs_ciphers` (Attributes) (see [below for nested schema](#nestedatt--split_proxy--server_non_pfs_ciphers))
<a id="nestedatt--split_proxy--mode"></a>
### Nested Schema for `split_proxy.mode`

Optional:

* `enable` (Boolean)
<a id="nestedatt--split_proxy--server_non_pfs_ciphers"></a>
### Nested Schema for `split_proxy.server_non_pfs_ciphers`

Optional:

* `enable` (Boolean)
<a id="nestedatt--start_tls"></a>
### Nested Schema for `start_tls`

Optional:

* `l4_port` (List of Number)
<a id="nestedatt--tcp"></a>
### Nested Schema for `tcp`

Optional:

* `delayed_ack` (Boolean) - enable/disable TCP delayed ACK
* `syn_retries` (Number) - TCP Sync retries
* `timewait_timeout` (Number) - TCP Wait Timeout value
<a id="nestedatt--tool"></a>
### Nested Schema for `tool`

Optional:

* `early_engage` (Boolean) - enable/disable tool early engage
* `fail_action` (String) - Action to take if the tool fails
<a id="nestedatt--tool_l3"></a>
### Nested Schema for `tool_l3`

Optional:

* `cache_server_cert_timeout` (Number) - cache server timeout in seconds.
* `http2_downgrade` (Attributes) (see [below for nested schema](#nestedatt--tool_l3--http2_downgrade))
* `nat_pat` (Attributes) (see [below for nested schema](#nestedatt--tool_l3--nat_pat))
<a id="nestedatt--tool_l3--http2_downgrade"></a>
### Nested Schema for `tool_l3.http2_downgrade`

Optional:

* `enable` (Boolean)
<a id="nestedatt--tool_l3--nat_pat"></a>
### Nested Schema for `tool_l3.nat_pat`

Optional:

* `enable` (Boolean)
<a id="nestedatt--url_cache"></a>
### Nested Schema for `url_cache`

Optional:

* `miss_action` (String) - The action to take if local URL category resolution misses
* `timeout` (Number) - defer timeout in seconds. Only applicable for missAction 'defer'

