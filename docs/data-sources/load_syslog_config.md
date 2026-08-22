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
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_name` (String, computed) - Name of the cluster
* `syslog_config_list` (List(Object({device_ip, log_severity, target_hosts})), computed) - List of the syslog configuration specification for every node in the cluster

