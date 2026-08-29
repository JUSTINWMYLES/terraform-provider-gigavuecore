---
page_title: "gigavuecore_load_ntp_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load NTP Configuration
---

# gigavuecore_load_ntp_config Data Source

Load NTP Configuration

## Example Usage

```terraform
data "gigavuecore_load_ntp_config" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auth_enabled` (Boolean, computed)
* `auth_keys` (Attributes List, computed) (see [below for nested schema](#nestedatt--auth_keys))
* `clock_sync` (Boolean, computed) - indicates system's clock is synchronized with referenced NTP server
* `enabled` (Boolean, computed) - enable/disable use of NTP for synchronization of the system's clock
* `ntp_servers` (Attributes List, computed) (see [below for nested schema](#nestedatt--ntp_servers))
* `ref_server` (String, computed) - Address of NTP server used to synchronize the system's clock. ipv4 or ipv6 or hostname
* `statuses` (Attributes List, computed) (see [below for nested schema](#nestedatt--statuses))

<a id="nestedatt--auth_keys"></a>
### Nested Schema for `auth_keys`

Read-Only:

* `key` (String) - MD5 key
* `key_number` (Number)
* `trusted` (Boolean)
<a id="nestedatt--ntp_servers"></a>
### Nested Schema for `ntp_servers`

Read-Only:

* `enabled` (Boolean)
* `key_enabled` (Boolean)
* `key_number` (Number)
* `preferred` (Boolean)
* `server` (String) - ipv4 or ipv6 or hostname
* `version` (String)
<a id="nestedatt--statuses"></a>
### Nested Schema for `statuses`

Read-Only:

* `address` (String) - ip address of the server
* `last_resp` (Number) - seconds
* `offset` (Number) - milliseconds in float
* `poll_interval` (Number) - seconds
* `ref_clock` (String)
* `status` (String) - server status
* `stratum` (Number)

