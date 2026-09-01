---
page_title: "gigavuecore_configure_acme_certificate_details_of_fm Action - gigavuecore"
subcategory: ""
description: |-
  renew/revoke ACME certificate of FM
---

# gigavuecore_configure_acme_certificate_details_of_fm Action

renew/revoke ACME certificate of FM

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_configure_acme_certificate_details_of_fm" "example" {
  config {
    acme_certificate = [{
      domain = "example"
    }]
    operation_type = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `acme_certificate` (Attributes List, optional) (see [below for nested schema](#nestedatt--acme_certificate))
* `operation_type` (String, required) - operationType of the acme certificate renew/revoke

<a id="nestedatt--acme_certificate"></a>
### Nested Schema for `acme_certificate`

Required:

* `domain` (String) - domain name can be FM IP or FQDN

