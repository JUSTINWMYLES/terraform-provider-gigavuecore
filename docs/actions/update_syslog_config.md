---
page_title: "gigavuecore_update_syslog_config Action - gigavuecore"
subcategory: ""
description: |-
  Update Syslog Config
---

# gigavuecore_update_syslog_config Action

Update Syslog Config

## Example Usage

```terraform
action "gigavuecore_update_syslog_config" "example" {
  config {
    cluster_id   = "example"
    cluster_name = "example"
    syslog_config_list = [{
      device_ip    = "example"
      log_severity = "example"
      target_hosts = [{
        log_severity      = "example"
        port              = 0
        server            = "example"
        ssh_enabled       = true
        streaming_enabled = true
        username          = "example"
      }]
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `cluster_name` (String, optional) - Name of the cluster
* `syslog_config_list` (Attributes List, optional) - List of the syslog configuration specification for every node in the cluster (see [below for nested schema](#nestedatt--syslog_config_list))

<a id="nestedatt--syslog_config_list"></a>
### Nested Schema for `syslog_config_list`

Required:

* `device_ip` (String) - IP address of the device
* `log_severity` (String) - Minimum syslog logging severity level

Optional:

* `target_hosts` (Attributes List) - List of syslog targets the device is streaming to (see [below for nested schema](#nestedatt--syslog_config_list--target_hosts))

<a id="nestedatt--syslog_config_list--target_hosts"></a>
### Nested Schema for `syslog_config_list.target_hosts`

Required:

* `log_severity` (String) - Minimum logging severity level the device will stream for this target
* `port` (Number) - 0 represents UDP and non zero is TCP
* `server` (String) - ipv4 or ipv6 or hostname
* `streaming_enabled` (Boolean) - Is Syslog Streaming enabled

Optional:

* `ssh_enabled` (Boolean) - Is Syslog target configured to receive logs via SSH
* `username` (String) - For syslog over UDP there won't be any user. Only valid for SSH type

