---
page_title: "gigavuecore_delete_acmecertificate_details_of_device Action - gigavuecore"
subcategory: ""
description: |-
  Delete ACME certificate details
---

# gigavuecore_delete_acmecertificate_details_of_device Action

Delete ACME certificate details

## Example Usage

```terraform
action "gigavuecore_delete_acmecertificate_details_of_device" "example" {
  config {
    acme_certificate = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `acme_certificate` (List of Dynamic, optional)


