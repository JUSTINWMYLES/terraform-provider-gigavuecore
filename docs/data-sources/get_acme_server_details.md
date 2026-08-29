---
page_title: "gigavuecore_get_acme_server_details Data Source - gigavuecore"
subcategory: ""
description: |-
  get a list of acme server details
---

# gigavuecore_get_acme_server_details Data Source

get a list of acme server details

## Example Usage

```terraform
data "gigavuecore_get_acme_server_details" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `acme_enabled` (Boolean)
* `acme_server_url` (String)
* `alias` (String)
* `file_name` (String)

