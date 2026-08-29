---
page_title: "gigavuecore_issue_acme_certificate Action - gigavuecore"
subcategory: ""
description: |-
  issue ACME certificate for FM
---

# gigavuecore_issue_acme_certificate Action

issue ACME certificate for FM

## Example Usage

```terraform
action "gigavuecore_issue_acme_certificate" "example" {
  config {
    acme_certificate = [{
      acme_server_alias = "example"
      algorithm         = "example"
      domain            = "example"
      renew_days        = 0
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

* `acme_server_alias` (String)
* `domain` (String) - domain name can be FM IP or FQDN
Optional:

* `algorithm` (String)
* `renew_days` (Number) - default will be 1/3rd of certificate validity period

