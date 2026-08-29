---
page_title: "gigavuecore_get_upgrade_info_by_task_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Upgrade Info By TaskId
---

# gigavuecore_get_upgrade_info_by_task_id Data Source

Get Upgrade Info By TaskId

## Example Usage

```terraform
data "gigavuecore_get_upgrade_info_by_task_id" "example" {
  task_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `task_id` (String, required)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `cluster_id` (String) - Cluster ID
* `cluster_log` (Attributes List) (see [below for nested schema](#nestedatt--items--cluster_log))
* `end_time` (String) - End Time
* `node_gs_upgrade_status` (Attributes List) (see [below for nested schema](#nestedatt--items--node_gs_upgrade_status))
* `node_upgrade_status` (Attributes List) (see [below for nested schema](#nestedatt--items--node_upgrade_status))
* `object_counts` (Attributes) (see [below for nested schema](#nestedatt--items--object_counts))
* `object_counts_expected` (Attributes) (see [below for nested schema](#nestedatt--items--object_counts_expected))
* `object_diff_report` (Attributes) (see [below for nested schema](#nestedatt--items--object_diff_report))
* `operation_state` (String) - Operation State
* `overall_upgrade_status` (String) - Overall Upgrade Status
* `stage` (String) - Stage
* `start_time` (String) - Start Time
* `status` (String) - Status
* `task_id` (String) - Task ID
* `task_name` (String) - Task Name
* `upgrade_flow` (String) - Upgrade Flow
* `upgrade_summary` (Attributes) (see [below for nested schema](#nestedatt--items--upgrade_summary))
<a id="nestedatt--items--cluster_log"></a>
### Nested Schema for `items.cluster_log`

Read-Only:

* `cluster_id` (String) - Cluster ID
* `message` (String) - Message
* `node_id` (String) - Node ID
* `time` (String) - Time
<a id="nestedatt--items--node_gs_upgrade_status"></a>
### Nested Schema for `items.node_gs_upgrade_status`

Read-Only:

* `current_version` (Dynamic) - Current Version
* `host_name` (String) - Host Name
* `node_id` (String) - Node ID
* `node_upgrade_state` (String) - Node Upgrade State
* `previous_version` (Dynamic) - Previous Version
* `product_code` (String) - Product Code
* `stack_mode` (String) - Stack Mode
* `step_status` (Attributes List) (see [below for nested schema](#nestedatt--items--node_gs_upgrade_status--step_status))
<a id="nestedatt--items--node_gs_upgrade_status--step_status"></a>
### Nested Schema for `items.node_gs_upgrade_status.step_status`

Read-Only:

* `message` (String) - Message
* `time` (String) - Time
<a id="nestedatt--items--node_upgrade_status"></a>
### Nested Schema for `items.node_upgrade_status`

Read-Only:

* `current_step` (Number) - Current Step
* `current_version` (String) - Current Version
* `file_name` (String) - File Name
* `file_server` (String) - File Server
* `host_name` (String) - Host Name
* `last_completed_stage` (String) - Last Completed Stage
* `node_id` (String) - Node ID
* `node_log` (Attributes List) (see [below for nested schema](#nestedatt--items--node_upgrade_status--node_log))
* `num_steps_in_task` (Number) - Number of Steps In Task
* `prior_version` (String) - Prior Version
* `status` (String) - Status
* `step_status` (String) - Step Status
<a id="nestedatt--items--node_upgrade_status--node_log"></a>
### Nested Schema for `items.node_upgrade_status.node_log`

Read-Only:

* `cluster_id` (String) - Cluster ID
* `message` (String) - Message
* `node_id` (String) - Node ID
* `time` (String) - Time
<a id="nestedatt--items--object_counts"></a>
### Nested Schema for `items.object_counts`

Read-Only:

* `cards_up` (Number) - Number of Cards Up
* `circuit_ports_up` (Number) - Number of Circuit Ports Up
* `gigasmart_ports_up` (Number) - Number of Giga Smart Ports Up
* `gsops` (Number) - Number Gsops
* `hybrid_ports_up` (Number) - Number of Hybrid Ports Up
* `inline_network_ports_up` (Number) - Number of Inline Network Ports Up
* `inline_tool_ports_up` (Number) - Number of Inline Tool Ports Up
* `maps` (Number) - Number of Maps
* `network_ports_up` (Number) - Number of Network Ports Up
* `nodes` (Number) - Number of Nodes
* `stack_links` (Number) - Number of Stack Links
* `stack_ports_up` (Number) - Number of Stack Ports Up
* `tool_ports_up` (Number) - Number of Tool Ports Up
<a id="nestedatt--items--object_counts_expected"></a>
### Nested Schema for `items.object_counts_expected`

Read-Only:

* `cards_up` (Number) - Number of Cards Up
* `circuit_ports_up` (Number) - Number of Circuit Ports Up
* `gigasmart_ports_up` (Number) - Number of Giga Smart Ports Up
* `gsops` (Number) - Number Gsops
* `hybrid_ports_up` (Number) - Number of Hybrid Ports Up
* `inline_network_ports_up` (Number) - Number of Inline Network Ports Up
* `inline_tool_ports_up` (Number) - Number of Inline Tool Ports Up
* `maps` (Number) - Number of Maps
* `network_ports_up` (Number) - Number of Network Ports Up
* `nodes` (Number) - Number of Nodes
* `stack_links` (Number) - Number of Stack Links
* `stack_ports_up` (Number) - Number of Stack Ports Up
* `tool_ports_up` (Number) - Number of Tool Ports Up
<a id="nestedatt--items--object_diff_report"></a>
### Nested Schema for `items.object_diff_report`

Read-Only:

* `cards_up` (Attributes List) - Cards Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--cards_up))
* `circuit_ports_up` (Attributes List) - Circuit Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--circuit_ports_up))
* `gigasmart_ports_up` (Attributes List) - Giga Smart Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--gigasmart_ports_up))
* `gsops` (Attributes List) - Gsops Differences (see [below for nested schema](#nestedatt--items--object_diff_report--gsops))
* `hybrid_ports_up` (Attributes List) - Hybrid Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--hybrid_ports_up))
* `inline_network_ports_up` (Attributes List) - Inline Network Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--inline_network_ports_up))
* `inline_tool_ports_up` (Attributes List) - Inline Tool Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--inline_tool_ports_up))
* `maps` (Attributes List) - Maps Differences (see [below for nested schema](#nestedatt--items--object_diff_report--maps))
* `network_ports_up` (Attributes List) - Network Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--network_ports_up))
* `nodes` (Attributes List) - Nodes Differences (see [below for nested schema](#nestedatt--items--object_diff_report--nodes))
* `stack_links` (Attributes List) - Stack Links Differences (see [below for nested schema](#nestedatt--items--object_diff_report--stack_links))
* `stack_ports_up` (Attributes List) - Stack Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--stack_ports_up))
* `tool_ports_up` (Attributes List) - Tool Ports Up Differences (see [below for nested schema](#nestedatt--items--object_diff_report--tool_ports_up))
<a id="nestedatt--items--object_diff_report--cards_up"></a>
### Nested Schema for `items.object_diff_report.cards_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--circuit_ports_up"></a>
### Nested Schema for `items.object_diff_report.circuit_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--gigasmart_ports_up"></a>
### Nested Schema for `items.object_diff_report.gigasmart_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--gsops"></a>
### Nested Schema for `items.object_diff_report.gsops`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--hybrid_ports_up"></a>
### Nested Schema for `items.object_diff_report.hybrid_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--inline_network_ports_up"></a>
### Nested Schema for `items.object_diff_report.inline_network_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--inline_tool_ports_up"></a>
### Nested Schema for `items.object_diff_report.inline_tool_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--maps"></a>
### Nested Schema for `items.object_diff_report.maps`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--network_ports_up"></a>
### Nested Schema for `items.object_diff_report.network_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--nodes"></a>
### Nested Schema for `items.object_diff_report.nodes`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--stack_links"></a>
### Nested Schema for `items.object_diff_report.stack_links`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--stack_ports_up"></a>
### Nested Schema for `items.object_diff_report.stack_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--object_diff_report--tool_ports_up"></a>
### Nested Schema for `items.object_diff_report.tool_ports_up`

Read-Only:

* `after_state` (String) - After State
* `object_name` (String) - Object Name
* `prev_state` (String) - Previous State
<a id="nestedatt--items--upgrade_summary"></a>
### Nested Schema for `items.upgrade_summary`

Read-Only:

* `num_activation_preparation_complete` (Number) - Activation Preparation Complete Count
* `num_config_backup_complete` (Number) - Config Backup Complete Count
* `num_init_fetch_complete` (Number) - Init Fetch Complete Count
* `num_initial_validation_complete` (Number) - CC-Card Sync Complete Count
* `num_install_complete` (Number) - Install Complete Count
* `num_nodes_success` (Number) - Nodes Success Count
* `num_post_upgrade_validation_complete` (Number) - Post Upgrade Validation Complete Count
* `num_reload_complete` (Number) - Reload Complete Count
* `num_reload_started` (Number) - Reload Started Count
* `num_uboot_complete` (Number) - UBoot Complete Count
* `num_upgrade_complete` (Number) - Upgrade Complete Count
* `num_verification_complete` (Number) - Verification Complete Count

