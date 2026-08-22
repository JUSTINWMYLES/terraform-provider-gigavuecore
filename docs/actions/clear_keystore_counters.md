---
page_title: "gigavuecore_clear_keystore_counters Action - gigavuecore"
subcategory: ""
description: |-
  Clear keystore hit counters
---

# gigavuecore_clear_keystore_counters Action

Clear keystore hit counters

## Example Usage

```terraform
action "gigavuecore_clear_keystore_counters" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
