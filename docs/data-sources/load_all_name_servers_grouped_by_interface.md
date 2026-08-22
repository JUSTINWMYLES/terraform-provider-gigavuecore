---
page_title: "gigavuecore_load_all_name_servers_grouped_by_interface Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all name servers grouped by Network Interface
---

# gigavuecore_load_all_name_servers_grouped_by_interface Data Source

Load all name servers grouped by Network Interface

## Example Usage

```terraform
data "gigavuecore_load_all_name_servers_grouped_by_interface" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({interface_name, servers})), computed)

