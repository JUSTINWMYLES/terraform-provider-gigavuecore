---
page_title: "gigavuecore_redefine_ntp_server Action - gigavuecore"
subcategory: ""
description: |-
  Redefine NTP Server configuration
---

# gigavuecore_redefine_ntp_server Action

Redefine NTP Server configuration

## Example Usage

```terraform
action "gigavuecore_redefine_ntp_server" "example" {
  config {
    cluster_id = "example"
    enabled = true
    key_enabled = true
    key_number = 1
    preferred = true
    server = "example"
    server_address = "example"
    version = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Bool, optional)
* `key_enabled` (Bool, optional)
* `key_number` (Number, optional)
* `preferred` (Bool, optional)
* `server` (String, required) - ipv4 or ipv6 or hostname
* `server_address` (String, required) - address of NTP Server to redefine
* `version` (String, optional)
