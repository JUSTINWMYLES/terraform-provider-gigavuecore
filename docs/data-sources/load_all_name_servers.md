---
page_title: "gigavuecore_load_all_name_servers Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all name servers
---

# gigavuecore_load_all_name_servers Data Source

Load all name servers

## Example Usage

```terraform
data "gigavuecore_load_all_name_servers" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({interface_name, servers})), computed)

