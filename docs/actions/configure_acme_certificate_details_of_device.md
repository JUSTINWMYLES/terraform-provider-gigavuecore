---
page_title: "gigavuecore_configure_acme_certificate_details_of_device Action - gigavuecore"
subcategory: ""
description: |-
  issue renew revoke ACME certificate of device
---

# gigavuecore_configure_acme_certificate_details_of_device Action

issue renew revoke ACME certificate of device

## Example Usage

```terraform
action "gigavuecore_configure_acme_certificate_details_of_device" "example" {
  config {
    acme_certificate = [{
      acme_server_url = "example"
      algorithm       = "example"
      domain          = "example"
      renew_days      = 0
    }]
    operation_type = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `acme_certificate` (Attributes List, optional) (see [below for nested schema](#nestedatt--acme_certificate))
* `operation_type` (String, required) - operationType of the acme certificate issue/renew/revoke

<a id="nestedatt--acme_certificate"></a>
### Nested Schema for `acme_certificate`

Required:

* `acme_server_url` (String)
* `domain` (String)
Optional:

* `algorithm` (String)
* `renew_days` (Number) - default will be 1/3rd of certificate validity period

