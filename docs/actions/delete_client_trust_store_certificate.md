---
page_title: "gigavuecore_delete_client_trust_store_certificate Action - gigavuecore"
subcategory: ""
description: |-
  Delete the Client Trust Store Certificate
---

# gigavuecore_delete_client_trust_store_certificate Action

Delete the Client Trust Store Certificate

## Example Usage

```terraform
action "gigavuecore_delete_client_trust_store_certificate" "example" {
  config {
    alias       = "example"
    cluster_id  = "example"
    fingerprint = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Client Trust Store alias
* `cluster_id` (String, required) - Target Cluster ID
* `fingerprint` (String, required) - fingerprint for the certificate


