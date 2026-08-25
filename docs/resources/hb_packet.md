---
page_title: "gigavuecore_hb_packet Resource - gigavuecore"
subcategory: ""
description: |-
  Find Heartbeat Packet by alias (deprecated: use GET /inline/hbProfiles/{alias})
---

# gigavuecore_hb_packet Resource

Find Heartbeat Packet by alias (deprecated: use GET /inline/hbProfiles/{alias})

## Example Usage

```terraform
resource "gigavuecore_hb_packet" "example" {
  alias         = null
  custom_packet = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Heartbeat Custom Packet alias. Maps directly (foreign key) into InlineHbProfile alias
* `custom_packet` (String, required) - Base64-encoded custom pcap packet


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hb_packet.example {alias}
```
