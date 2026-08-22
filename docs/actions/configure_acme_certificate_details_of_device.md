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
    acme_certificate = null
    operation_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `acme_certificate` (List(Dynamic), optional)
* `operation_type` (String, required) - operationType of the acme certificate issue/renew/revoke
