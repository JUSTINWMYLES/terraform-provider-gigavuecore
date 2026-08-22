---
page_title: "gigavuecore_remove_hash_tool_mapping Action - gigavuecore"
subcategory: ""
description: |-
  Remove mapping of hash bucket to port
---

# gigavuecore_remove_hash_tool_mapping Action

Remove mapping of hash bucket to port

## Example Usage

```terraform
action "gigavuecore_remove_hash_tool_mapping" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    hash_bucket_ids = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target gigastream
* `cluster_id` (String, required) - Target Cluster ID
* `hash_bucket_ids` (String, required) - id or range of hash bucket to clear the port mapping
