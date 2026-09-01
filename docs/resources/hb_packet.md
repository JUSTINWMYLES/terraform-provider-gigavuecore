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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hb_packet.example {alias}/{cluster_id}
```
