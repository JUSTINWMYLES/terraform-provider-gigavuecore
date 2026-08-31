---
page_title: "gigavuecore_upload_heartbeat_packet_from_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload a Heartbeat Packet from local file
---

# gigavuecore_upload_heartbeat_packet_from_file Action

Upload a Heartbeat Packet from local file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_upload_heartbeat_packet_from_file" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    hb_packet  = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Heartbeat Profile
* `cluster_id` (String, required) - Target Cluster ID
* `hb_packet` (String, required) - Attached Heartbeat Packet. In native binary format


