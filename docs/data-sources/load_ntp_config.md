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
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `auth_enabled` (Bool, computed)
* `auth_keys` (List(Object({key, key_number, trusted})), computed)
* `clock_sync` (Bool, computed) - indicates system's clock is synchronized with referenced NTP server
* `enabled` (Bool, computed) - enable/disable use of NTP for synchronization of the system's clock
* `ntp_servers` (List(Object({enabled, key_enabled, key_number, preferred, server, version})), computed)
* `ref_server` (String, computed) - Address of NTP server used to synchronize the system's clock. ipv4 or ipv6 or hostname
* `statuses` (List(Object({address, last_resp, offset, poll_interval, ref_clock, status, stratum})), computed)

