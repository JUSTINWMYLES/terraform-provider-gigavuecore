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
  cluster_id = null
  page       = null
  sort       = null
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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `inline_ssl_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--inline_ssl_profiles))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--inline_ssl_profiles"></a>
### Nested Schema for `inline_ssl_profiles`

Read-Only:

* `alias` (String) - Alias of the inline SSL profile
* `certificate` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--certificate))
* `cluster_id` (String) - id of the defining cluster
* `decrypt` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--decrypt))
* `default_action` (String) - Action to take if none of the profile rules match
* `high_avail` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--high_avail))
* `inbound_tool_early_inspect` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--inbound_tool_early_inspect))
* `key_map` (Attributes List) (see [below for nested schema](#nestedatt--inline_ssl_profiles--key_map))
* `monitor` (String)
* `network_group` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--network_group))
* `no_decrypt` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--no_decrypt))
* `non_ssl_tcp` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--non_ssl_tcp))
* `one_arm` (String)
* `resilient_inline` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--resilient_inline))
* `rules` (List of Dynamic) - inline SSL profile rules
* `split_proxy` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--split_proxy))
* `start_tls` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--start_tls))
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--tcp))
* `tool` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--tool))
* `tool_l3` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--tool_l3))
* `url_cache` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--url_cache))
<a id="nestedatt--inline_ssl_profiles--certificate"></a>
### Nested Schema for `inline_ssl_profiles.certificate`

Read-Only:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--certificate--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate
<a id="nestedatt--inline_ssl_profiles--certificate--revocation"></a>
### Nested Schema for `inline_ssl_profiles.certificate.revocation`

Read-Only:

* `crl` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--certificate--revocation--crl))
* `ocsp` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--certificate--revocation--ocsp))
<a id="nestedatt--inline_ssl_profiles--certificate--revocation--crl"></a>
### Nested Schema for `inline_ssl_profiles.certificate.revocation.crl`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--inline_ssl_profiles--certificate--revocation--ocsp"></a>
### Nested Schema for `inline_ssl_profiles.certificate.revocation.ocsp`

Read-Only:

* `defer` (Number) - timeout in seconds
* `enabled` (Boolean)
* `fail` (String) - only applicable when enabled
<a id="nestedatt--inline_ssl_profiles--decrypt"></a>
### Nested Schema for `inline_ssl_profiles.decrypt`

Read-Only:

* `tcp` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--decrypt--tcp))
* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--decrypt--tool_bypass))
<a id="nestedatt--inline_ssl_profiles--decrypt--tcp"></a>
### Nested Schema for `inline_ssl_profiles.decrypt.tcp`

Read-Only:

* `inactive_timeout` (Number) - SSL decryption TCP inactive timeout (in minutes)
* `port_map` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--decrypt--tcp--port_map))
<a id="nestedatt--inline_ssl_profiles--decrypt--tcp--port_map"></a>
### Nested Schema for `inline_ssl_profiles.decrypt.tcp.port_map`

Read-Only:

* `default_out_port` (Number) - egress port for decryption port map. 0 is disabled.
* `ports` (Attributes List) (see [below for nested schema](#nestedatt--inline_ssl_profiles--decrypt--tcp--port_map--ports))
<a id="nestedatt--inline_ssl_profiles--decrypt--tcp--port_map--ports"></a>
### Nested Schema for `inline_ssl_profiles.decrypt.tcp.port_map.ports`

Read-Only:

* `in_port` (Number) - ingress port for decryption port map
* `out_port` (Number) - egress port for decryption port map
* `rule_id` (Number)
<a id="nestedatt--inline_ssl_profiles--decrypt--tool_bypass"></a>
### Nested Schema for `inline_ssl_profiles.decrypt.tool_bypass`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--high_avail"></a>
### Nested Schema for `inline_ssl_profiles.high_avail`

Read-Only:

* `active_standby` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--high_avail--active_standby))
<a id="nestedatt--inline_ssl_profiles--high_avail--active_standby"></a>
### Nested Schema for `inline_ssl_profiles.high_avail.active_standby`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--inbound_tool_early_inspect"></a>
### Nested Schema for `inline_ssl_profiles.inbound_tool_early_inspect`

Read-Only:

* `connection_timeout` (Number) - connection timeout timeout in seconds.
* `mode` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--inbound_tool_early_inspect--mode))
<a id="nestedatt--inline_ssl_profiles--inbound_tool_early_inspect--mode"></a>
### Nested Schema for `inline_ssl_profiles.inbound_tool_early_inspect.mode`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--key_map"></a>
### Nested Schema for `inline_ssl_profiles.key_map`

Read-Only:

* `hostname` (String) - hostname or IP address
* `key` (String) - SSL key alias
* `rule_id` (Number)
<a id="nestedatt--inline_ssl_profiles--network_group"></a>
### Nested Schema for `inline_ssl_profiles.network_group`

Read-Only:

* `multiple_entry` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--network_group--multiple_entry))
<a id="nestedatt--inline_ssl_profiles--network_group--multiple_entry"></a>
### Nested Schema for `inline_ssl_profiles.network_group.multiple_entry`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--no_decrypt"></a>
### Nested Schema for `inline_ssl_profiles.no_decrypt`

Read-Only:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--no_decrypt--tool_bypass))
<a id="nestedatt--inline_ssl_profiles--no_decrypt--tool_bypass"></a>
### Nested Schema for `inline_ssl_profiles.no_decrypt.tool_bypass`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--non_ssl_tcp"></a>
### Nested Schema for `inline_ssl_profiles.non_ssl_tcp`

Read-Only:

* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--non_ssl_tcp--tool_bypass))
<a id="nestedatt--inline_ssl_profiles--non_ssl_tcp--tool_bypass"></a>
### Nested Schema for `inline_ssl_profiles.non_ssl_tcp.tool_bypass`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--resilient_inline"></a>
### Nested Schema for `inline_ssl_profiles.resilient_inline`

Read-Only:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--resilient_inline--mode))
<a id="nestedatt--inline_ssl_profiles--resilient_inline--mode"></a>
### Nested Schema for `inline_ssl_profiles.resilient_inline.mode`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--split_proxy"></a>
### Nested Schema for `inline_ssl_profiles.split_proxy`

Read-Only:

* `mode` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--split_proxy--mode))
* `server_non_pfs_ciphers` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--split_proxy--server_non_pfs_ciphers))
<a id="nestedatt--inline_ssl_profiles--split_proxy--mode"></a>
### Nested Schema for `inline_ssl_profiles.split_proxy.mode`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--split_proxy--server_non_pfs_ciphers"></a>
### Nested Schema for `inline_ssl_profiles.split_proxy.server_non_pfs_ciphers`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--start_tls"></a>
### Nested Schema for `inline_ssl_profiles.start_tls`

Read-Only:

* `l4_port` (List of Number)
<a id="nestedatt--inline_ssl_profiles--tcp"></a>
### Nested Schema for `inline_ssl_profiles.tcp`

Read-Only:

* `delayed_ack` (Boolean) - enable/disable TCP delayed ACK
* `syn_retries` (Number) - TCP Sync retries
* `timewait_timeout` (Number) - TCP Wait Timeout value
<a id="nestedatt--inline_ssl_profiles--tool"></a>
### Nested Schema for `inline_ssl_profiles.tool`

Read-Only:

* `early_engage` (Boolean) - enable/disable tool early engage
* `fail_action` (String) - Action to take if the tool fails
<a id="nestedatt--inline_ssl_profiles--tool_l3"></a>
### Nested Schema for `inline_ssl_profiles.tool_l3`

Read-Only:

* `cache_server_cert_timeout` (Number) - cache server timeout in seconds.
* `http2_downgrade` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--tool_l3--http2_downgrade))
* `nat_pat` (Attributes) (see [below for nested schema](#nestedatt--inline_ssl_profiles--tool_l3--nat_pat))
<a id="nestedatt--inline_ssl_profiles--tool_l3--http2_downgrade"></a>
### Nested Schema for `inline_ssl_profiles.tool_l3.http2_downgrade`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--tool_l3--nat_pat"></a>
### Nested Schema for `inline_ssl_profiles.tool_l3.nat_pat`

Read-Only:

* `enable` (Boolean)
<a id="nestedatt--inline_ssl_profiles--url_cache"></a>
### Nested Schema for `inline_ssl_profiles.url_cache`

Read-Only:

* `miss_action` (String) - The action to take if local URL category resolution misses
* `timeout` (Number) - defer timeout in seconds. Only applicable for missAction 'defer'

