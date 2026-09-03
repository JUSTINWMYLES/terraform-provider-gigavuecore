---
page_title: "gigavuecore_notif_config Resource - gigavuecore"
subcategory: ""
description: |-
  Modify Notification Config
---

# gigavuecore_notif_config Resource

Modify Notification Config

~> **Note:** This resource is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
resource "gigavuecore_notif_config" "example" {
  interface_name = "example"
  interface_type = "management"
  target_address = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `interface_name` (String, optional) - Name of any one of the available network interface names on the FM. If this is chosen, FM registers itself as a notification target on the node with Ipv6 address if both FM and the node have Ipv6 address otherwise Ipv4 address is used. If targetAddress is configured then interfaceName has no effect.
* `interface_type` (String, required) - Notification target interface type
* `target_address` (String, optional) - Configure FM's DNS name or static IP address to receive the management or data traffic from the node. The configured address is used by FM to register itself as a notification target on the node

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

