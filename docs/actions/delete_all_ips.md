---
page_title: "gigavuecore_delete_all_ips Action - gigavuecore"
subcategory: ""
description: |-
  new in H 6.8
---

# gigavuecore_delete_all_ips Action

new in H 6.8

## Example Usage

```terraform
action "gigavuecore_delete_all_ips" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    ip_interface = "example"
    port = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the gpfcp profile
* `cluster_id` (String, required) - Target Cluster ID
* `ip_interface` (String, required) - IP Interface address
* `port` (String, required) - Port or Port List separated by commas
