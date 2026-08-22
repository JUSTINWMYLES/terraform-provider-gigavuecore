---
page_title: "gigavuecore_get_all_external_export_target_server Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all External Export Server
---

# gigavuecore_get_all_external_export_target_server Data Source

Load all External Export Server

## Example Usage

```terraform
data "gigavuecore_get_all_external_export_target_server" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, auth_type, broker_type, host, password, port, user_name})), computed)

