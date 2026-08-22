---
page_title: "gigavuecore_notif_config Resource - gigavuecore"
subcategory: ""
description: |-
  Load Fm Notification Target Config By Interface Type
---

# gigavuecore_notif_config Resource

Load Fm Notification Target Config By Interface Type

## Example Usage

```terraform
resource "gigavuecore_notif_config" "example" {
  id = null
  interface_name = null
  interface_type = null
  target_address = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `id` (String, required)
* `interface_name` (String, optional) - Name of any one of the available network interface names on the FM. If this is chosen, FM registers itself as a notification target on the node with Ipv6 address if both FM and the node have Ipv6 address otherwise Ipv4 address is used. If targetAddress is configured then interfaceName has no effect.
* `interface_type` (String, required) - Notification target interface type
* `target_address` (String, optional) - Configure FM's DNS name or static IP address to receive the management or data traffic from the node. The configured address is used by FM to register itself as a notification target on the node

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `interface_name` (String, computed) - Name of any one of the available network interface names on the FM. If this is chosen, FM registers itself as a notification target on the node with Ipv6 address if both FM and the node have Ipv6 address otherwise Ipv4 address is used. If targetAddress is configured then interfaceName has no effect.
* `target_address` (String, computed) - Configure FM's DNS name or static IP address to receive the management or data traffic from the node. The configured address is used by FM to register itself as a notification target on the node

