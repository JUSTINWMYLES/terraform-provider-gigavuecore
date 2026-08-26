---
page_title: "gigavuecore_delete_ntp_auth_key Action - gigavuecore"
subcategory: ""
description: |-
  Delete NTP Auth Key by alias
---

# gigavuecore_delete_ntp_auth_key Action

Delete NTP Auth Key by alias

## Example Usage

```terraform
action "gigavuecore_delete_ntp_auth_key" "example" {
  config {
    auth_key_id = 1
    cluster_id  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `auth_key_id` (Number, required) - NTP Authentication Key Id
* `cluster_id` (String, required) - Target Cluster ID


