---
page_title: "gigavuecore_delete_all_icap_servers Action - gigavuecore"
subcategory: ""
description: |-
  Delete all ICAP Servers
---

# gigavuecore_delete_all_icap_servers Action

Delete all ICAP Servers

## Example Usage

```terraform
action "gigavuecore_delete_all_icap_servers" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


