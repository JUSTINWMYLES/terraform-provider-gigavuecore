---
page_title: "gigavuecore_reset_ssl_trust_store Action - gigavuecore"
subcategory: ""
description: |-
  Reset inline SSL trust-store to default trust-store
---

# gigavuecore_reset_ssl_trust_store Action

Reset inline SSL trust-store to default trust-store

## Example Usage

```terraform
action "gigavuecore_reset_ssl_trust_store" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


