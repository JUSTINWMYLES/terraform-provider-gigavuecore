---
page_title: "gigavuecore_get_system_hostname Data Source - gigavuecore"
subcategory: ""
description: |-
  get system hostname configurations
---

# gigavuecore_get_system_hostname Data Source

get system hostname configurations

## Example Usage

```terraform
data "gigavuecore_get_system_hostname" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Gigamon cluster Id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alternate_hostname_with_dhcp` (String, computed) - indicates alternate host name with DHCP client request
* `hostname` (String, computed) - hostname of system
* `hostname_with_dhcp` (Boolean, computed) - enable to use host name with DHCP client request


