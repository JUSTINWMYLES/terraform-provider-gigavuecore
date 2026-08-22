---
page_title: "gigavuecore_load_all_system_acme_certificate_details Data Source - gigavuecore"
subcategory: ""
description: |-
  get a list of deviceacme certificate details
---

# gigavuecore_load_all_system_acme_certificate_details Data Source

get a list of deviceacme certificate details

## Example Usage

```terraform
data "gigavuecore_load_all_system_acme_certificate_details" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({box_id, cluster_name, issued_cert, last_acme_request})), computed)

