---
page_title: "gigavuecore_load_all_inline_ssl_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all inline SSL profiles
---

# gigavuecore_load_all_inline_ssl_profiles Data Source

Load all inline SSL profiles

## Example Usage

```terraform
data "gigavuecore_load_all_inline_ssl_profiles" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias of the inline SSL profile
* `certificate` (Attributes) (see [below for nested schema](#nestedatt--items--certificate))
* `cluster_id` (String) - id of the defining cluster
* `decrypt` (Attributes) (see [below for nested schema](#nestedatt--items--decrypt))
* `default_action` (String) - Action to take if none of the profile rules match
* `high_avail` (Attributes) (see [below for nested schema](#nestedatt--items--high_avail))
* `inbound_tool_early_inspect` (Attributes) (see [below for nested schema](#nestedatt--items--inbound_tool_early_inspect))
* `key_map` (Attributes List) (see [below for nested schema](#nestedatt--items--key_map))
* `monitor` (String)
* `network_group` (Attributes) (see [below for nested schema](#nestedatt--items--network_group))
* `no_decrypt` (Attributes) (see [below for nested schema](#nestedatt--items--no_decrypt))
* `non_ssl_tcp` (Attributes) (see [below for nested schema](#nestedatt--items--non_ssl_tcp))
* `one_arm` (String)
* `resilient_inline` (Attributes) (see [below for nested schema](#nestedatt--items--resilient_inline))
* `rules` (List of Dynamic) - inline SSL profile rules
* `split_proxy` (Attributes) (see [below for nested schema](#nestedatt--items--split_proxy))
* `start_tls` (Attributes) (see [below for nested schema](#nestedatt--items--start_tls))
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--items--tcp))
* `tool` (Attributes) (see [below for nested schema](#nestedatt--items--tool))
* `tool_l3` (Attributes) (see [below for nested schema](#nestedatt--items--tool_l3))
* `url_cache` (Attributes) (see [below for nested schema](#nestedatt--items--url_cache))
<a id="nestedatt--items--certificate"></a>
### Nested Schema for `items.certificate`

Read-Only:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) (see [below for nested schema](#nestedatt--items--certificate--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate
<a id="nestedatt--items--certificate--revocation"></a>
### Nested Schema for `items.certificate.revocation`

Read-Only:

* `crl` (Attributes) (see [below for nested schema](#nestedatt--items--certificate--revocation--crl))
* `ocsp` (Attributes) (see [below for nested schema](#nestedatt--items--certificate--revocation--ocsp))
<a id="nestedatt--items--certificate--revocation--crl"></a>
### Nested Schema for `items.certificate.revocation.crl`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--items--certificate--revocation--ocsp"></a>
### Nested Schema for `items.certificate.revocation.ocsp`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--items--decrypt"></a>
### Nested Schema for `items.decrypt`

Read-Only:

* `tcp` (Attributes) (see [below for nested schema](#nestedatt--items--decrypt--tcp))
* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--items--decrypt--tool_bypass))
<a id="nestedatt--items--decrypt--tcp"></a>
### Nested Schema for `items.decrypt.tcp`

Read-Only:

* `inactive_timeout` (Number) - SSL decryption TCP inactive timeout (in minutes)
* `port_map` (Attributes) (see [below for nested schema](#nestedatt--items--decrypt--tcp--port_map))
<a id="nestedatt--items--decrypt--tcp--port_map"></a>
### Nested Schema for `items.decrypt.tcp.port_map`

Read-Only:

* `default_out_port` (Number) - egress port for decryption port map. 0 is disabled.
* `ports` (Attributes List) (see [below for nested schema](#nestedatt--items--decrypt--tcp--port_map--ports))
<a id="nestedatt--items--decrypt--tcp--port_map--ports"></a>
### Nested Schema for `items.decrypt.tcp.port_map.ports`

Read-Only:

* `in_port` (Number) - ingress port for decryption port map
* `out_port` (Number) - egress port for decryption port map
* `rule_id` (Number)
<a id="nestedatt--items--decrypt--tool_bypass"></a>
### Nested Schema for `items.decrypt.tool_bypass`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--high_avail"></a>
### Nested Schema for `items.high_avail`

Read-Only:

* `active_standby` (Attributes) (see [below for nested schema](#nestedatt--items--high_avail--active_standby))
<a id="nestedatt--items--high_avail--active_standby"></a>
### Nested Schema for `items.high_avail.active_standby`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--inbound_tool_early_inspect"></a>
### Nested Schema for `items.inbound_tool_early_inspect`

Read-Only:

* `connection_timeout` (Number) - connection timeout timeout in seconds.
* `mode` (Attributes) (see [below for nested schema](#nestedatt--items--inbound_tool_early_inspect--mode))
<a id="nestedatt--items--inbound_tool_early_inspect--mode"></a>
### Nested Schema for `items.inbound_tool_early_inspect.mode`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--key_map"></a>
### Nested Schema for `items.key_map`

Read-Only:

* `hostname` (String) - hostname or IP address
* `key` (String) - SSL key alias
* `rule_id` (Number)
<a id="nestedatt--items--network_group"></a>
### Nested Schema for `items.network_group`

Read-Only:

* `multiple_entry` (Attributes) (see [below for nested schema](#nestedatt--items--network_group--multiple_entry))
<a id="nestedatt--items--network_group--multiple_entry"></a>
### Nested Schema for `items.network_group.multiple_entry`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--no_decrypt"></a>
### Nested Schema for `items.no_decrypt`

Read-Only:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--items--no_decrypt--tool_bypass))
<a id="nestedatt--items--no_decrypt--tool_bypass"></a>
### Nested Schema for `items.no_decrypt.tool_bypass`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--non_ssl_tcp"></a>
### Nested Schema for `items.non_ssl_tcp`

Read-Only:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--items--non_ssl_tcp--tool_bypass))
<a id="nestedatt--items--non_ssl_tcp--tool_bypass"></a>
### Nested Schema for `items.non_ssl_tcp.tool_bypass`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--resilient_inline"></a>
### Nested Schema for `items.resilient_inline`

Read-Only:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--items--resilient_inline--mode))
<a id="nestedatt--items--resilient_inline--mode"></a>
### Nested Schema for `items.resilient_inline.mode`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--split_proxy"></a>
### Nested Schema for `items.split_proxy`

Read-Only:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--items--split_proxy--mode))
* `server_non_pfs_ciphers` (Attributes) (see [below for nested schema](#nestedatt--items--split_proxy--server_non_pfs_ciphers))
<a id="nestedatt--items--split_proxy--mode"></a>
### Nested Schema for `items.split_proxy.mode`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--split_proxy--server_non_pfs_ciphers"></a>
### Nested Schema for `items.split_proxy.server_non_pfs_ciphers`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--start_tls"></a>
### Nested Schema for `items.start_tls`

Read-Only:

* `l4_port` (List of Number)
<a id="nestedatt--items--tcp"></a>
### Nested Schema for `items.tcp`

Read-Only:

* `delayed_ack` (Boolean) - enable/disable TCP delayed ACK
* `syn_retries` (Number) - TCP Sync retries
* `timewait_timeout` (Number) - TCP Wait Timeout value
<a id="nestedatt--items--tool"></a>
### Nested Schema for `items.tool`

Read-Only:

* `early_engage` (Boolean) - enable/disable tool early engage
* `fail_action` (String) - Action to take if the tool fails
<a id="nestedatt--items--tool_l3"></a>
### Nested Schema for `items.tool_l3`

Read-Only:

* `cache_server_cert_timeout` (Number) - cache server timeout in seconds.
* `http2_downgrade` (Attributes) (see [below for nested schema](#nestedatt--items--tool_l3--http2_downgrade))
* `nat_pat` (Attributes) (see [below for nested schema](#nestedatt--items--tool_l3--nat_pat))
<a id="nestedatt--items--tool_l3--http2_downgrade"></a>
### Nested Schema for `items.tool_l3.http2_downgrade`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--tool_l3--nat_pat"></a>
### Nested Schema for `items.tool_l3.nat_pat`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--items--url_cache"></a>
### Nested Schema for `items.url_cache`

Read-Only:

* `miss_action` (String) - The action to take if local URL category resolution misses
* `timeout` (Number) - defer timeout in seconds. Only applicable for missAction 'defer'

