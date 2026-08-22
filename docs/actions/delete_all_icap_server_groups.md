---
page_title: "gigavuecore_delete_all_icap_server_groups Action - gigavuecore"
subcategory: ""
description: |-
  Delete all ICAP Server Groups
---

# gigavuecore_delete_all_icap_server_groups Action

Delete all ICAP Server Groups

## Example Usage

```terraform
action "gigavuecore_delete_all_icap_server_groups" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
