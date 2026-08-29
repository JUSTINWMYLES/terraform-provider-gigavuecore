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
  alias = "example"
  certificate = {
    expired = "example"
    invalid = "example"
    revocation = {
      crl = {
        defer   = 0
        enabled = true
        fail    = "example"
      }
      ocsp = {
        defer   = 0
        enabled = true
        fail    = "example"
      }
    }
    self_signed = "example"
    unknown_ca  = "example"
  }
  cluster_id = "example"
  decrypt = {
    tcp = {
      inactive_timeout = 0
      port_map = {
        default_out_port = 0
        ports = [{
          in_port  = 0
          out_port = 0
          rule_id  = 0
        }]
      }
    }
    tool_bypass = {
      enable = true
    }
  }
  default_action = "example"
  high_avail = {
    active_standby = {
      enable = true
    }
  }
  inbound_tool_early_inspect = {
    connection_timeout = 0
    mode = {
      enable = true
    }
  }
  key_map = [{
    hostname = "example"
    key      = "example"
    rule_id  = 0
  }]
  monitor = "example"
  network_group = {
    multiple_entry = {
      enable = true
    }
  }
  no_decrypt = {
    tool_bypass = {
      enable = true
    }
  }
  non_ssl_tcp = {
    tool_bypass = {
      enable = true
    }
  }
  one_arm = "example"
  resilient_inline = {
    mode = {
      enable = true
    }
  }
  rules = [ "example" ]
  split_proxy = {
    mode = {
      enable = true
    }
    server_non_pfs_ciphers = {
      enable = true
    }
  }
  start_tls = {
    l4_port = [ 0 ]
  }
  tcp = {
    delayed_ack      = true
    syn_retries      = 0
    timewait_timeout = 0
  }
  tool = {
    early_engage = true
    fail_action  = "example"
  }
  tool_l3 = {
    cache_server_cert_timeout = 0
    http2_downgrade = {
      enable = true
    }
    nat_pat = {
      enable = true
    }
  }
  url_cache = {
    miss_action = "example"
    timeout     = 0
  }
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `certificate` (Attributes, optional) - SSL profile certificate configuration (see [below for nested schema](#nestedatt--certificate))
* `cluster_id` (String, required) - id of the defining cluster
* `decrypt` (Attributes, optional) - SSL profile configuration on decrypt action (see [below for nested schema](#nestedatt--decrypt))
* `default_action` (String, optional) - Action to take if none of the profile rules match
* `high_avail` (Attributes, optional) - SSL profile configuration on high availability (see [below for nested schema](#nestedatt--high_avail))
* `inbound_tool_early_inspect` (Attributes, optional) - SSL profile configuration for InboundToolEarlyInspect (see [below for nested schema](#nestedatt--inbound_tool_early_inspect))
* `key_map` (Attributes List, optional) (see [below for nested schema](#nestedatt--key_map))
* `monitor` (String, optional)
* `network_group` (Attributes, optional) - SSL Profile configuration for multiple entry in network groups (see [below for nested schema](#nestedatt--network_group))
* `no_decrypt` (Attributes, optional) - SSL profile configuration on no-decrypt action (see [below for nested schema](#nestedatt--no_decrypt))
* `non_ssl_tcp` (Attributes, optional) - SSL profile configuration on TCP proxy action (see [below for nested schema](#nestedatt--non_ssl_tcp))
* `one_arm` (String, optional)
* `resilient_inline` (Attributes, optional) - SSL profile configuration for Resilient Inline Arrangement (see [below for nested schema](#nestedatt--resilient_inline))
* `rules` (List of Dynamic, optional) - inline SSL profile rules
* `split_proxy` (Attributes, optional) - SSL profile configuration for Split Proxy (see [below for nested schema](#nestedatt--split_proxy))
* `start_tls` (Attributes, optional) - SSL profile configuration on start TLS action (see [below for nested schema](#nestedatt--start_tls))
* `tcp` (Attributes, optional) - SSL profile configuration on TCP (see [below for nested schema](#nestedatt--tcp))
* `tool` (Attributes, optional) - SSL Profile configuration for Tools (see [below for nested schema](#nestedatt--tool))
* `tool_l3` (Attributes, optional) - SSL profile configuration for Layer 3 ISSL (see [below for nested schema](#nestedatt--tool_l3))
* `url_cache` (Attributes, optional) - SSL profile configuration on url-cache (see [below for nested schema](#nestedatt--url_cache))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--certificate"></a>
### Nested Schema for `certificate`

Optional:

* `expired` (String) - SSL profile on expired certificate
* `invalid` (String) - SSL profile on invalid certificate
* `revocation` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--certificate--revocation))
* `self_signed` (String) - SSL profile on self-signed certificate
* `unknown_ca` (String) - SSL profile on unknown CA certificate

<a id="nestedatt--certificate--revocation"></a>
### Nested Schema for `certificate.revocation`

Optional:

* `crl` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--certificate--revocation--crl))
* `ocsp` (Attributes) - SSL profile certificate revocation configuration (see [below for nested schema](#nestedatt--certificate--revocation--ocsp))

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

* `tcp` (Attributes) - SSL decryption TCP control (see [below for nested schema](#nestedatt--decrypt--tcp))
* `tool_bypass` (Attributes) (see [below for nested schema](#nestedatt--decrypt--tool_bypass))

<a id="nestedatt--decrypt--tcp"></a>
### Nested Schema for `decrypt.tcp`

Optional:

* `inactive_timeout` (Number) - SSL decryption TCP inactive timeout (in minutes)
* `port_map` (Attributes) - SSL decryption port map (see [below for nested schema](#nestedatt--decrypt--tcp--port_map))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_inline_ssl_profile.example {alias}:{cluster_id}
```
