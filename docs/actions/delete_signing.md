---
page_title: "gigavuecore_delete_signing Action - gigavuecore"
subcategory: ""
description: |-
  Delete the primary or secondary certificate and key
---

# gigavuecore_delete_signing Action

Delete the primary or secondary certificate and key

## Example Usage

```terraform
action "gigavuecore_delete_signing" "example" {
  config {
    cluster_id  = "example"
    signing_for = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `signing_for` (String, required)


