---
page_title: "gigavuecore_update_system_host_banner Action - gigavuecore"
subcategory: ""
description: |-
  Update system banner configurations
---

# gigavuecore_update_system_host_banner Action

Update system banner configurations

## Example Usage

```terraform
action "gigavuecore_update_system_host_banner" "example" {
  config {
    body_cluster_id       = "example"
    cluster_id            = "example"
    failed_login_attempts = "example"
    login_msg             = "example"
    motd                  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_cluster_id` (String, optional) - Gigamon cluster Id
* `cluster_id` (String, required) - Target Cluster ID
* `failed_login_attempts` (String, optional) - failed LoginAttempts
* `login_msg` (String, optional) - Login Message
* `motd` (String, optional) - Message Of The Day


