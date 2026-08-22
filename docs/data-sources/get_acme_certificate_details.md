---
page_title: "gigavuecore_get_acme_certificate_details Data Source - gigavuecore"
subcategory: ""
description: |-
  get a list of acme certificateDetails
---

# gigavuecore_get_acme_certificate_details Data Source

get a list of acme certificateDetails

## Example Usage

```terraform
data "gigavuecore_get_acme_certificate_details" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({cert_last_request_status, domain, issued_cert_details, task_id})), computed)

