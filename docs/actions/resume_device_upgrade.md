---
page_title: "gigavuecore_resume_device_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Resume Device Upgrade
---

# gigavuecore_resume_device_upgrade Action

Resume Device Upgrade

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_resume_device_upgrade" "example" {
  config {
    device_upgrade_resume_specs = [{
      activate_stage   = true
      cluster_name     = "example"
      config_backup    = true
      device_ip        = "example"
      fetch_stage      = true
      install_stage    = true
      post_check_stage = true
      task_id          = "example"
    }]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `device_upgrade_resume_specs` (Attributes List, required) (see [below for nested schema](#nestedatt--device_upgrade_resume_specs))

<a id="nestedatt--device_upgrade_resume_specs"></a>
### Nested Schema for `device_upgrade_resume_specs`

Required:

* `activate_stage` (Boolean) - Activate Stage
* `cluster_name` (String) - Cluster Name
* `config_backup` (Boolean) - Config Backup
* `device_ip` (String) - Device IP
* `fetch_stage` (Boolean) - Fetch Stage
* `install_stage` (Boolean) - Install Stage
* `post_check_stage` (Boolean) - Post Check Stage
* `task_id` (String) - Task ID

