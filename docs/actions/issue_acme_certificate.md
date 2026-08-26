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
    acme_certificate = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `acme_certificate` (List of Dynamic, optional)


