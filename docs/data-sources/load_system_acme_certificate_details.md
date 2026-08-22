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
* `issued_cert` (Object({acme_ca_url, acme_service, algorithm, cert_name, domain, expiry, first_issued, last_failed_renew, last_success_renew, next_renew, renew_days, status}), computed)
  * `acme_ca_url` (String, computed)
  * `acme_service` (String, computed)
  * `algorithm` (String, computed)
  * `cert_name` (String, computed)
  * `domain` (String, computed)
  * `expiry` (String, computed)
  * `first_issued` (String, computed)
  * `last_failed_renew` (String, computed)
  * `last_success_renew` (String, computed)
  * `next_renew` (String, computed)
  * `renew_days` (Number, computed) - default will be 1/3rd of certificate validity period
  * `status` (String, computed)
* `last_acme_request` (Object({acme_ca_url, domain_name, status, type}), computed)
  * `acme_ca_url` (String, computed)
  * `domain_name` (String, computed)
  * `status` (String, computed)
  * `type` (String, computed)

