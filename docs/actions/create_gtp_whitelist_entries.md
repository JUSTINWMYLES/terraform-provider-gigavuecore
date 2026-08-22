---
page_title: "gigavuecore_create_gtp_whitelist_entries Action - gigavuecore"
subcategory: ""
description: |-
  Create GTP Whitelist Entries
---

# gigavuecore_create_gtp_whitelist_entries Action

Create GTP Whitelist Entries

## Example Usage

```terraform
action "gigavuecore_create_gtp_whitelist_entries" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    entries = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID
* `entries` (List(Dynamic), required)
