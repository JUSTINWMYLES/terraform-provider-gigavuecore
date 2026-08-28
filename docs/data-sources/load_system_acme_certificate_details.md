---
page_title: "gigavuecore_load_system_acme_certificate_details Data Source - gigavuecore"
subcategory: ""
description: |-
  get deviceacme certificate details
---

# gigavuecore_load_system_acme_certificate_details Data Source

get deviceacme certificate details

## Example Usage

```terraform
data "gigavuecore_load_system_acme_certificate_details" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `box_id` (String, computed)
* `cluster_name` (String, computed)
* `issued_cert` (Attributes, computed) (see [below for nested schema](#nestedatt--issued_cert))
* `last_acme_request` (Attributes, computed) (see [below for nested schema](#nestedatt--last_acme_request))

<a id="nestedatt--issued_cert"></a>
### Nested Schema for `issued_cert`

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
<a id="nestedatt--last_acme_request"></a>
### Nested Schema for `last_acme_request`

Read-Only:

* `acme_ca_url` (String)
* `domain_name` (String)
* `status` (String)
* `type` (String)

