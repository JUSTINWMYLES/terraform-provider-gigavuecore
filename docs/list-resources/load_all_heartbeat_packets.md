---
page_title: "gigavuecore_load_all_heartbeat_packets List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Heartbeat Packets (deprecated: use GET /inline/hbProfiles)
---

# gigavuecore_load_all_heartbeat_packets List Resource

Load all Heartbeat Packets (deprecated: use GET /inline/hbProfiles)

## Example Usage

```terraform
list "gigavuecore_load_all_heartbeat_packets" "example" {
  provider = gigavuecore
  limit = 100
}

```

## Schema

### Arguments

The following arguments are supported:

