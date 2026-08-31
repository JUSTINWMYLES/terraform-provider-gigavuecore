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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `interface_name` (String) - Interface in which the nameserver resides
* `servers` (List of String) - List of name server address

