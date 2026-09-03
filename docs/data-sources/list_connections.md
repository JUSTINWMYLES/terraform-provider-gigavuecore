---
page_title: "gigavuecore_list_connections Data Source - gigavuecore"
subcategory: ""
description: |-
  List all unified resource connections by environment id
---

# gigavuecore_list_connections Data Source

List all unified resource connections by environment id

## Example Usage

```terraform
data "gigavuecore_list_connections" "example" {
  env_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `unified_resources_connection_list_query_response` (Dynamic, computed) - Unified Resource connections


