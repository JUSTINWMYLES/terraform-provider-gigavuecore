---
page_title: "gigavuecore_add_ntp_auth_key Action - gigavuecore"
subcategory: ""
description: |-
  Add an NTP Authentication Key
---

# gigavuecore_add_ntp_auth_key Action

Add an NTP Authentication Key

## Example Usage

```terraform
action "gigavuecore_add_ntp_auth_key" "example" {
  config {
    cluster_id = "example"
    key = "example"
    key_number = 1
    trusted = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `key` (String, optional) - MD5 key
* `key_number` (Number, required)
* `trusted` (Bool, optional)
