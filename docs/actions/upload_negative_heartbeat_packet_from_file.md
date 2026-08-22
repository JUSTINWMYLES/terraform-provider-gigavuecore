---
page_title: "gigavuecore_upload_negative_heartbeat_packet_from_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload a Negative Heartbeat Packet from local file
---

# gigavuecore_upload_negative_heartbeat_packet_from_file Action

Upload a Negative Heartbeat Packet from local file

## Example Usage

```terraform
action "gigavuecore_upload_negative_heartbeat_packet_from_file" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    nhb_packet = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Negative Heartbeat Profile
* `cluster_id` (String, required) - Target Cluster ID
* `nhb_packet` (String, required) - Attached Negative Heartbeat Packet. In native binary format
