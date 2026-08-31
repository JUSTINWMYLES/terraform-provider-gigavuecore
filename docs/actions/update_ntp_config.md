---
page_title: "gigavuecore_update_ntp_config Action - gigavuecore"
subcategory: ""
description: |-
  Update NTP Configuration
---

# gigavuecore_update_ntp_config Action

Update NTP Configuration

## Example Usage

```terraform
action "gigavuecore_update_ntp_config" "example" {
  config {
    auth_enabled = true
    auth_keys = [{
      key        = "example"
      key_number = 1
      trusted    = true
    }]
    clock_sync = true
    cluster_id = "example"
    enabled    = true
    ntp_servers = [{
      enabled     = true
      key_enabled = true
      key_number  = 1
      preferred   = true
      server      = "example"
      version     = "v3"
    }]
    ref_server = "example"
    statuses = [{
      address       = "example"
      last_resp     = 0
      offset        = 1.0
      poll_interval = 0
      ref_clock     = "example"
      status        = "example"
      stratum       = 0
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `auth_enabled` (Boolean, optional)
* `auth_keys` (Attributes List, optional) (see [below for nested schema](#nestedatt--auth_keys))
* `clock_sync` (Boolean, optional) - indicates system's clock is synchronized with referenced NTP server
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional) - enable/disable use of NTP for synchronization of the system's clock
* `ntp_servers` (Attributes List, optional) (see [below for nested schema](#nestedatt--ntp_servers))
* `ref_server` (String, optional) - Address of NTP server used to synchronize the system's clock. ipv4 or ipv6 or hostname
* `statuses` (Attributes List, optional) (see [below for nested schema](#nestedatt--statuses))

<a id="nestedatt--auth_keys"></a>
### Nested Schema for `auth_keys`

Required:

* `key_number` (Number)

Optional:

* `key` (String) - MD5 key
* `trusted` (Boolean)

<a id="nestedatt--ntp_servers"></a>
### Nested Schema for `ntp_servers`

Required:

* `server` (String) - ipv4 or ipv6 or hostname

Optional:

* `enabled` (Boolean)
* `key_enabled` (Boolean)
* `key_number` (Number)
* `preferred` (Boolean)
* `version` (String)

<a id="nestedatt--statuses"></a>
### Nested Schema for `statuses`

Optional:

* `address` (String) - ip address of the server
* `last_resp` (Number) - seconds
* `offset` (Number) - milliseconds in float
* `poll_interval` (Number) - seconds
* `ref_clock` (String)
* `status` (String) - server status
* `stratum` (Number)

