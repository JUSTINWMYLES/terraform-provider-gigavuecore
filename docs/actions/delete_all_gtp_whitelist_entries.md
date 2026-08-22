---
page_title: "gigavuecore_delete_all_gtp_whitelist_entries Action - gigavuecore"
subcategory: ""
description: |-
  Delete All GTP Whitelist Entries
---

# gigavuecore_delete_all_gtp_whitelist_entries Action

Delete All GTP Whitelist Entries

## Example Usage

```terraform
action "gigavuecore_delete_all_gtp_whitelist_entries" "example" {
  config {
    alias = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID
