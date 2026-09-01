---
page_title: "gigavuecore_update_system_web_configuration Action - gigavuecore"
subcategory: ""
description: |-
  Update system web configuration
---

# gigavuecore_update_system_web_configuration Action

Update system web configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_system_web_configuration" "example" {
  config {
    auto_logout_timeout    = 0
    cluster_id             = "example"
    cluster_name           = "example"
    enable_https           = true
    http_port              = 0
    https_port             = 0
    server_ssl_min_version = "tls1"
    session_renewal        = 0
    session_timeout        = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `auto_logout_timeout` (Number, required) - Auto Logout Timeout in seconds
* `cluster_id` (String, optional) - Target Cluster Id
* `cluster_name` (String, optional) - cluster name
* `enable_https` (Boolean, required) - enableHttps
* `http_port` (Number, required) - Http Port Number
* `https_port` (Number, required) - Https Port Number
* `server_ssl_min_version` (String, optional) - SSL/TLS minimum version supported
* `session_renewal` (Number, required) - Session renewal in seconds
* `session_timeout` (Number, required) - Session renewal in seconds


