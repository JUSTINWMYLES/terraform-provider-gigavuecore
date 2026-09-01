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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias for external export target server
* `auth_type` (String) - Auth Type for external export target server
* `broker_type` (String) - Broker Type for external export target server
* `host` (String) - Host for external export target server
* `password` (String) - Password for external export target server
* `port` (String) - Port for external export target server
* `user_name` (String) - Username for external export target server

