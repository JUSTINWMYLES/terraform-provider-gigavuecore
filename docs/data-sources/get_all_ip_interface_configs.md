---
page_title: "gigavuecore_get_all_ip_interface_configs Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all ipInterface solution configurations
---

# gigavuecore_get_all_ip_interface_configs Data Source

Load all ipInterface solution configurations

## Example Usage

```terraform
data "gigavuecore_get_all_ip_interface_configs" "example" {
  application = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `application` (String, optional) - Gets ipInterface solution that supports the specified GigaSmart application
* `cluster_id` (String, optional) - Gets ipInterface solution configs for the provided cluster

### Attributes

In addition to all arguments above, the following attributes are exported:

* `ip_interface_configs` (List(Object({alias, applications, cluster_name, config_status, config_status_reasons, gateway, interfaces, ip_address, ip_mask, managed_status, mtu, ref_count})), computed)
* `tags` (List(Object({tag_key, tag_values})), computed) - RBAC Tags

