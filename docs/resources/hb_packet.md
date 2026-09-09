---
page_title: "gigavuecore_hb_packet Resource - gigavuecore"
subcategory: ""
description: |-
  Upload a new Heartbeat Packet (deprecated: use PATCH /inline/hbProfiles/{alias})
---

# gigavuecore_hb_packet Resource

Upload a new Heartbeat Packet (deprecated: use PATCH /inline/hbProfiles/{alias})

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

## Example Usage

```terraform
resource "gigavuecore_hb_packet" "example" {
  alias         = "example"
  cluster_id    = "example"
  custom_packet = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Heartbeat Custom Packet alias. Maps directly (foreign key) into InlineHbProfile alias
* `cluster_id` (String, required) - Target Cluster ID
* `custom_packet` (String, required) - Base64-encoded custom pcap packet

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hb_packet.example {alias}/{cluster_id}
```
