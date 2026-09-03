---
page_title: "gigavuecore_remove_certificate_from_ca_list Action - gigavuecore"
subcategory: ""
description: |-
  delete a certificate from the list of trusted certificate authorities (CA List)
---

# gigavuecore_remove_certificate_from_ca_list Action

delete a certificate from the list of trusted certificate authorities (CA List)

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


