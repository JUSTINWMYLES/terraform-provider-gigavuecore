---
page_title: "gigavuecore_delete_all_icap_profiles Action - gigavuecore"
subcategory: ""
description: |-
  Delete all ICAP Profiles
---

# gigavuecore_delete_all_icap_profiles Action

Delete all ICAP Profiles

## Example Usage

```terraform
action "gigavuecore_delete_all_icap_profiles" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


