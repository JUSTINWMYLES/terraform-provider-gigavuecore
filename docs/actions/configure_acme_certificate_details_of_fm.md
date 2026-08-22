---
page_title: "gigavuecore_configure_acme_certificate_details_of_fm Action - gigavuecore"
subcategory: ""
description: |-
  renew/revoke ACME certificate of FM
---

# gigavuecore_configure_acme_certificate_details_of_fm Action

renew/revoke ACME certificate of FM

## Example Usage

```terraform
action "gigavuecore_configure_acme_certificate_details_of_fm" "example" {
  config {
    acme_certificate = null
    operation_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `acme_certificate` (List(Dynamic), optional)
* `operation_type` (String, required) - operationType of the acme certificate renew/revoke
