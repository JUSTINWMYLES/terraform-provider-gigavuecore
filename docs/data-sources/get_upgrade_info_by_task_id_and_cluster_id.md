---
page_title: "gigavuecore_get_upgrade_info_by_task_id_and_cluster_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Upgrade Info By TaskId and ClusterId
---

# gigavuecore_get_upgrade_info_by_task_id_and_cluster_id Data Source

Get Upgrade Info By TaskId and ClusterId

## Example Usage

```terraform
data "gigavuecore_get_upgrade_info_by_task_id_and_cluster_id" "example" {
  cluster_id = null
  task_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster ID
* `task_id` (String, required) - Task ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_log` (List(Object({cluster_id, message, node_id, time})), computed)
* `end_time` (String, computed) - End Time
* `node_gs_upgrade_status` (List(Object({current_version, host_name, node_id, node_upgrade_state, previous_version, product_code, stack_mode, step_status})), computed)
* `node_upgrade_status` (List(Object({current_step, current_version, file_name, file_server, host_name, last_completed_stage, node_id, node_log, num_steps_in_task, prior_version, status, step_status})), computed)
* `object_counts` (Object({cards_up, circuit_ports_up, gigasmart_ports_up, gsops, hybrid_ports_up, inline_network_ports_up, inline_tool_ports_up, maps, network_ports_up, nodes, stack_links, stack_ports_up, tool_ports_up}), computed)
  * `cards_up` (Number, computed) - Number of Cards Up
  * `circuit_ports_up` (Number, computed) - Number of Circuit Ports Up
  * `gigasmart_ports_up` (Number, computed) - Number of Giga Smart Ports Up
  * `gsops` (Number, computed) - Number Gsops
  * `hybrid_ports_up` (Number, computed) - Number of Hybrid Ports Up
  * `inline_network_ports_up` (Number, computed) - Number of Inline Network Ports Up
  * `inline_tool_ports_up` (Number, computed) - Number of Inline Tool Ports Up
  * `maps` (Number, computed) - Number of Maps
  * `network_ports_up` (Number, computed) - Number of Network Ports Up
  * `nodes` (Number, computed) - Number of Nodes
  * `stack_links` (Number, computed) - Number of Stack Links
  * `stack_ports_up` (Number, computed) - Number of Stack Ports Up
  * `tool_ports_up` (Number, computed) - Number of Tool Ports Up
* `object_counts_expected` (Object({cards_up, circuit_ports_up, gigasmart_ports_up, gsops, hybrid_ports_up, inline_network_ports_up, inline_tool_ports_up, maps, network_ports_up, nodes, stack_links, stack_ports_up, tool_ports_up}), computed)
  * `cards_up` (Number, computed) - Number of Cards Up
  * `circuit_ports_up` (Number, computed) - Number of Circuit Ports Up
  * `gigasmart_ports_up` (Number, computed) - Number of Giga Smart Ports Up
  * `gsops` (Number, computed) - Number Gsops
  * `hybrid_ports_up` (Number, computed) - Number of Hybrid Ports Up
  * `inline_network_ports_up` (Number, computed) - Number of Inline Network Ports Up
  * `inline_tool_ports_up` (Number, computed) - Number of Inline Tool Ports Up
  * `maps` (Number, computed) - Number of Maps
  * `network_ports_up` (Number, computed) - Number of Network Ports Up
  * `nodes` (Number, computed) - Number of Nodes
  * `stack_links` (Number, computed) - Number of Stack Links
  * `stack_ports_up` (Number, computed) - Number of Stack Ports Up
  * `tool_ports_up` (Number, computed) - Number of Tool Ports Up
* `object_diff_report` (Object({cards_up, circuit_ports_up, gigasmart_ports_up, gsops, hybrid_ports_up, inline_network_ports_up, inline_tool_ports_up, maps, network_ports_up, nodes, stack_links, stack_ports_up, tool_ports_up}), computed)
  * `cards_up` (List(Object({after_state, object_name, prev_state})), computed) - Cards Up Differences
  * `circuit_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Circuit Ports Up Differences
  * `gigasmart_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Giga Smart Ports Up Differences
  * `gsops` (List(Object({after_state, object_name, prev_state})), computed) - Gsops Differences
  * `hybrid_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Hybrid Ports Up Differences
  * `inline_network_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Inline Network Ports Up Differences
  * `inline_tool_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Inline Tool Ports Up Differences
  * `maps` (List(Object({after_state, object_name, prev_state})), computed) - Maps Differences
  * `network_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Network Ports Up Differences
  * `nodes` (List(Object({after_state, object_name, prev_state})), computed) - Nodes Differences
  * `stack_links` (List(Object({after_state, object_name, prev_state})), computed) - Stack Links Differences
  * `stack_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Stack Ports Up Differences
  * `tool_ports_up` (List(Object({after_state, object_name, prev_state})), computed) - Tool Ports Up Differences
* `operation_state` (String, computed) - Operation State
* `overall_upgrade_status` (String, computed) - Overall Upgrade Status
* `stage` (String, computed) - Stage
* `start_time` (String, computed) - Start Time
* `status` (String, computed) - Status
* `task_name` (String, computed) - Task Name
* `upgrade_flow` (String, computed) - Upgrade Flow
* `upgrade_summary` (Object({num_activation_preparation_complete, num_config_backup_complete, num_init_fetch_complete, num_initial_validation_complete, num_install_complete, num_nodes_success, num_post_upgrade_validation_complete, num_reload_complete, num_reload_started, num_uboot_complete, num_upgrade_complete, num_verification_complete}), computed)
  * `num_activation_preparation_complete` (Number, computed) - Activation Preparation Complete Count
  * `num_config_backup_complete` (Number, computed) - Config Backup Complete Count
  * `num_init_fetch_complete` (Number, computed) - Init Fetch Complete Count
  * `num_initial_validation_complete` (Number, computed) - CC-Card Sync Complete Count
  * `num_install_complete` (Number, computed) - Install Complete Count
  * `num_nodes_success` (Number, computed) - Nodes Success Count
  * `num_post_upgrade_validation_complete` (Number, computed) - Post Upgrade Validation Complete Count
  * `num_reload_complete` (Number, computed) - Reload Complete Count
  * `num_reload_started` (Number, computed) - Reload Started Count
  * `num_uboot_complete` (Number, computed) - UBoot Complete Count
  * `num_upgrade_complete` (Number, computed) - Upgrade Complete Count
  * `num_verification_complete` (Number, computed) - Verification Complete Count

