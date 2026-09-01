---
page_title: "gigavuecore_delete_acmecertificate_details_of_device Action - gigavuecore"
subcategory: ""
description: |-
  Delete ACME certificate details
---

# gigavuecore_delete_acmecertificate_details_of_device Action

Delete ACME certificate details

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_acmecertificate_details_of_device" "example" {
  config {
    acme_certificate = [{
      domain = "example"
    }]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `acme_certificate` (Attributes List, optional) (see [below for nested schema](#nestedatt--acme_certificate))

<a id="nestedatt--acme_certificate"></a>
### Nested Schema for `acme_certificate`

Required:

* `domain` (String) - domain name can be FM IP or FQDN

