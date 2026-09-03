---
page_title: "gigavuecore_ntp_sync Action - gigavuecore"
subcategory: ""
description: |-
  Perform one-time synchronization of the H Series node's system clock with a specified NTP server.
---

# gigavuecore_ntp_sync Action

Perform one-time synchronization of the H Series node's system clock with a specified NTP server.

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


