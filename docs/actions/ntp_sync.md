---
page_title: "gigavuecore_ntp_sync Action - gigavuecore"
subcategory: ""
description: |-
  Perform one-time synchronization of the H Series node's system clock with a specified NTP server.
---

# gigavuecore_ntp_sync Action

Perform one-time synchronization of the H Series node's system clock with a specified NTP server.

## Example Usage

```terraform
action "gigavuecore_ntp_sync" "example" {
  config {
    cluster_id = "example"
    server     = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `server` (String, required) - ipv4 or ipv6 or hostname


