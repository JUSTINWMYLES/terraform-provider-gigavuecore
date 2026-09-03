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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cert_last_request_status` (Attributes) (see [below for nested schema](#nestedatt--items--cert_last_request_status))
* `domain` (String) - domain name will be FM IP or FQDN
* `issued_cert_details` (Attributes) (see [below for nested schema](#nestedatt--items--issued_cert_details))
* `task_id` (String)

<a id="nestedatt--items--cert_last_request_status"></a>
### Nested Schema for `items.cert_last_request_status`

Read-Only:

* `acme_server_alias` (String) - only for issue acmeServerAlias details will be shown
* `cert_status` (String)
* `domain` (String) - domain name can be FM IP or FQDN
* `request_type` (String)

<a id="nestedatt--items--issued_cert_details"></a>
### Nested Schema for `items.issued_cert_details`

Read-Only:

* `acme_server_alias` (String)
* `algorithm` (String)
* `expiry` (String)
* `first_issued` (String) - Date of the first issued certificate for FM
* `last_failed_renew` (String)
* `last_success_renew` (String)
* `next_renew` (String)
* `renew_days` (String) - only user Configured renewal days is shown
* `status` (String)
* `task_group_id` (String)

