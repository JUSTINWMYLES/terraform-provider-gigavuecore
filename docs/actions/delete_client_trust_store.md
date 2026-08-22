---
page_title: "gigavuecore_delete_client_trust_store Action - gigavuecore"
subcategory: ""
description: |-
  Delete the SSL Client Trust Store
---

# gigavuecore_delete_client_trust_store Action

Delete the SSL Client Trust Store

## Example Usage

```terraform
action "gigavuecore_delete_client_trust_store" "example" {
  config {
    alias = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Client Trust Store alias
* `cluster_id` (String, required) - Target Cluster ID
