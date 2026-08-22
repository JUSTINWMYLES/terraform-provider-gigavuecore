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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({acme_enabled, acme_server_url, alias, file_name})), computed)

