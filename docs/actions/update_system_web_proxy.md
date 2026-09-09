---
page_title: "gigavuecore_update_system_web_proxy Action - gigavuecore"
subcategory: ""
description: |-
  Update system web proxy
---

# gigavuecore_update_system_web_proxy Action

Update system web proxy

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_update_system_web_proxy" "example" {
  config {
    auth_type     = "none"
    cluster_id    = "example"
    cluster_name  = "example"
    password      = "example"
    proxy_address = "example"
    proxy_port    = 0
    username      = "example"
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


