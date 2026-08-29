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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (String)
* `cluster_name` (String)
* `issued_cert` (Attributes) (see [below for nested schema](#nestedatt--items--issued_cert))
* `last_acme_request` (Attributes) (see [below for nested schema](#nestedatt--items--last_acme_request))

<a id="nestedatt--items--issued_cert"></a>
### Nested Schema for `items.issued_cert`

Read-Only:

* `acme_ca_url` (String)
* `acme_service` (String)
* `algorithm` (String)
* `cert_name` (String)
* `domain` (String)
* `expiry` (String)
* `first_issued` (String)
* `last_failed_renew` (String)
* `last_success_renew` (String)
* `next_renew` (String)
* `renew_days` (Number) - default will be 1/3rd of certificate validity period
* `status` (String)

<a id="nestedatt--items--last_acme_request"></a>
### Nested Schema for `items.last_acme_request`

Read-Only:

* `acme_ca_url` (String)
* `domain_name` (String)
* `status` (String)
* `type` (String)

