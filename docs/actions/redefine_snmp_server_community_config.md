---
page_title: "gigavuecore_redefine_snmp_server_community_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Snmp Server Community config
---

# gigavuecore_redefine_snmp_server_community_config Action

Redefine Snmp Server Community config

## Example Usage

```terraform
action "gigavuecore_redefine_snmp_server_community_config" "example" {
  config {
    cluster_id = "example"
    community_strings = [ "example" ]
    enable_community_auth = true
    enable_community_auth_v1 = true
    enable_multi_community = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `community_strings` (List(String), optional) - community string\[s\] used to connect to this node using SNMP. The default value is 'public'. If 'enableMultiCommunity' is enabled, multiple community strings for the node are allowed
* `enable_community_auth` (Bool, optional) - turn on community-based authentication for the system
* `enable_community_auth_v1` (Bool, optional) - turn on community-based authentication for the SNMP v1
* `enable_multi_community` (Bool, optional) - allow configuration of multiple communities
