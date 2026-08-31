---
page_title: "gigavuecore_load_syslog_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Syslog config
---

# gigavuecore_load_syslog_config Data Source

Load Syslog config

## Example Usage

```terraform
data "gigavuecore_load_syslog_config" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_name` (String, computed) - Name of the cluster
* `syslog_config_list` (Attributes List, computed) - List of the syslog configuration specification for every node in the cluster (see [below for nested schema](#nestedatt--syslog_config_list))

<a id="nestedatt--syslog_config_list"></a>
### Nested Schema for `syslog_config_list`

Read-Only:

* `device_ip` (String) - IP address of the device
* `log_severity` (String) - Minimum syslog logging severity level
* `target_hosts` (Attributes List) - List of syslog targets the device is streaming to (see [below for nested schema](#nestedatt--syslog_config_list--target_hosts))

<a id="nestedatt--syslog_config_list--target_hosts"></a>
### Nested Schema for `syslog_config_list.target_hosts`

Read-Only:

* `log_severity` (String) - Minimum logging severity level the device will stream for this target
* `port` (Number) - 0 represents UDP and non zero is TCP
* `server` (String) - ipv4 or ipv6 or hostname
* `ssh_enabled` (Boolean) - Is Syslog target configured to receive logs via SSH
* `streaming_enabled` (Boolean) - Is Syslog Streaming enabled
* `username` (String) - For syslog over UDP there won't be any user. Only valid for SSH type

