---
page_title: "gigavuecore_update_system_web_proxy Action - gigavuecore"
subcategory: ""
description: |-
  Update system web proxy
---

# gigavuecore_update_system_web_proxy Action

Update system web proxy

## Example Usage

```terraform
action "gigavuecore_update_system_web_proxy" "example" {
  config {
    auth_type = "example"
    cluster_id = "example"
    cluster_name = "example"
    password = "example"
    proxy_address = "example"
    proxy_port = 1
    username = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `auth_type` (String, required) - Auth type
* `cluster_id` (String, optional) - Target Cluster Id
* `cluster_name` (String, optional) - cluster name
* `password` (String, required) - Proxy Password
* `proxy_address` (String, required) - Proxy address
* `proxy_port` (Number, required) - Proxy port
* `username` (String, required) - Proxy Username
