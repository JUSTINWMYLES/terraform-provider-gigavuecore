---
page_title: "gigavuecore_remove_certificate_from_ca_list Action - gigavuecore"
subcategory: ""
description: |-
  delete a certificate from the list of trusted certificate authorities (CA List)
---

# gigavuecore_remove_certificate_from_ca_list Action

delete a certificate from the list of trusted certificate authorities (CA List)

## Example Usage

```terraform
action "gigavuecore_remove_certificate_from_ca_list" "example" {
  config {
    cluster_id = "example"
    name       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `name` (String, required) - the certificate name


