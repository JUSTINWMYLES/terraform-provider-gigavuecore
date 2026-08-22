---
page_title: "gigavuecore_redefine_ntp_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine NTP Configuration
---

# gigavuecore_redefine_ntp_config Action

Redefine NTP Configuration

## Example Usage

```terraform
action "gigavuecore_redefine_ntp_config" "example" {
  config {
    auth_enabled = true
    auth_keys = null
    clock_sync = true
    cluster_id = "example"
    enabled = true
    ntp_servers = null
    ref_server = "example"
    statuses = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `auth_enabled` (Bool, optional)
* `auth_keys` (List(Dynamic), optional)
* `clock_sync` (Bool, optional) - indicates system's clock is synchronized with referenced NTP server
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Bool, optional) - enable/disable use of NTP for synchronization of the system's clock
* `ntp_servers` (List(Dynamic), optional)
* `ref_server` (String, optional) - Address of NTP server used to synchronize the system's clock. ipv4 or ipv6 or hostname
* `statuses` (List(Dynamic), optional)
