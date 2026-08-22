---
page_title: "gigavuecore_upload_heartbeat_packet_from_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload a Heartbeat Packet from local file
---

# gigavuecore_upload_heartbeat_packet_from_file Action

Upload a Heartbeat Packet from local file

## Example Usage

```terraform
action "gigavuecore_upload_heartbeat_packet_from_file" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    hb_packet = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Heartbeat Profile
* `cluster_id` (String, required) - Target Cluster ID
* `hb_packet` (String, required) - Attached Heartbeat Packet. In native binary format
