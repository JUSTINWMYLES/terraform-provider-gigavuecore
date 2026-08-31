---
page_title: "gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id Data Source - gigavuecore"
subcategory: ""
description: |-
  get the cluster configuration imageUpgrade status by taskGroupId
---

# gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id Data Source

get the cluster configuration imageUpgrade status by taskGroupId

## Example Usage

```terraform
data "gigavuecore_get_cluster_config_image_upgrade_status_by_task_group_id" "example" {
  task_group_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `task_group_id` (String, required) - system generated id

### Attributes

In addition to all arguments above, the following attributes are exported:

* `end_time` (String, computed) - End Time in UTC format
* `multi_upgrade_status` (Attributes List, computed) - list of the upgrade Status (see [below for nested schema](#nestedatt--multi_upgrade_status))
* `num_nodes` (Number, computed) - total number of nodes
* `num_nodes_success` (Number, computed) - total number of nodes success state
* `start_time` (String, computed) - Start Time in UTC format
* `status` (String, computed) - Image Upgrade Status
* `task_name` (String, computed) - name of the task

<a id="nestedatt--multi_upgrade_status"></a>
### Nested Schema for `multi_upgrade_status`

Read-Only:

* `cluster_id` (String) - Id of the Target cluster
* `cluster_log` (Attributes List) - list of the cluster Log (see [below for nested schema](#nestedatt--multi_upgrade_status--cluster_log))
* `end_time` (String) - End Time in UTC format
* `node_upgrade_status` (Attributes List) - list of the nodeUpgrade status (see [below for nested schema](#nestedatt--multi_upgrade_status--node_upgrade_status))
* `object_counts` (Attributes) - Cluster objectCounts (see [below for nested schema](#nestedatt--multi_upgrade_status--object_counts))
* `object_counts_expected` (Attributes) - Cluster objectCounts Expected (see [below for nested schema](#nestedatt--multi_upgrade_status--object_counts_expected))
* `object_diff_report` (Attributes) - Cluster object Diff report (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report))
* `start_time` (String) - Start Time in UTC format
* `status` (String) - Cluster Image Upgrade Status
* `task_id` (String) - taskId of the task
* `task_name` (String) - name of the task
* `upgrade_summary` (Attributes) - Cluster upgrade Summary (see [below for nested schema](#nestedatt--multi_upgrade_status--upgrade_summary))

<a id="nestedatt--multi_upgrade_status--cluster_log"></a>
### Nested Schema for `multi_upgrade_status.cluster_log`

Read-Only:

* `cluster_id` (String) - Id of the Target cluster
* `message` (String) - Read-only. Describes the reason for image upgrade state
* `node_id` (String) - unique ID representing device. This will be available only for nodeLog
* `time` (String) - In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'

<a id="nestedatt--multi_upgrade_status--node_upgrade_status"></a>
### Nested Schema for `multi_upgrade_status.node_upgrade_status`

Read-Only:

* `current_step` (Number) - current step number of task
* `current_version` (Number) - Current Device Version
* `file_name` (String) - Image filename
* `file_server` (String) - IpAddress of the Destination Image file server
* `hostname` (String) - hostname or IP address
* `node_id` (String) - unique ID representing device. This will be available in nodeLog
* `node_log` (Attributes List) - list of the node Log (see [below for nested schema](#nestedatt--multi_upgrade_status--node_upgrade_status--node_log))
* `num_steps_in_task` (Number) - total number of steps in task
* `prior_version` (Number) - Device Version before upgrade
* `status` (String) - node upgrade status
* `step_status` (String) - Describes the status of current step

<a id="nestedatt--multi_upgrade_status--node_upgrade_status--node_log"></a>
### Nested Schema for `multi_upgrade_status.node_upgrade_status.node_log`

Read-Only:

* `cluster_id` (String) - Id of the Target cluster
* `message` (String) - Read-only. Describes the reason for image upgrade state
* `node_id` (String) - unique ID representing device. This will be available only for nodeLog
* `time` (String) - In ISO-8601 date format 'yyyy-MM-dd'T'HH:mm:ssZ'

<a id="nestedatt--multi_upgrade_status--object_counts"></a>
### Nested Schema for `multi_upgrade_status.object_counts`

Read-Only:

* `cards_up` (Number) - total number of cards in up state
* `circuit_ports_up` (Number) - total number of circuit ports in up state
* `gigasmart_ports_up` (Number) - total number of gigasmart ports in up state
* `gsops` (Number) - total number of gigasmart operations
* `hybrid_ports_up` (Number) - total number of hybrid ports in up state
* `inline_network_ports_up` (Number) - total number of inlineNetwork ports in up state
* `inline_tool_ports_up` (Number) - total number of inlineTool ports in up state
* `maps` (Number) - total number of maps
* `network_ports_up` (Number) - total number of network ports in up state
* `nodes` (Number) - total number of nodes
* `stack_links` (Number) - total number of stacklinks
* `stack_ports_up` (Number) - total number of stack ports in up state
* `tool_ports_up` (Number) - total number of tool ports in up state

<a id="nestedatt--multi_upgrade_status--object_counts_expected"></a>
### Nested Schema for `multi_upgrade_status.object_counts_expected`

Read-Only:

* `cards_up` (Number) - total number of cards in up state
* `circuit_ports_up` (Number) - total number of circuit ports in up state
* `gigasmart_ports_up` (Number) - total number of gigasmart ports in up state
* `gsops` (Number) - total number of gigasmart operations
* `hybrid_ports_up` (Number) - total number of hybrid ports in up state
* `inline_network_ports_up` (Number) - total number of inlineNetwork ports in up state
* `inline_tool_ports_up` (Number) - total number of inlineTool ports in up state
* `maps` (Number) - total number of maps
* `network_ports_up` (Number) - total number of network ports in up state
* `nodes` (Number) - total number of nodes
* `stack_links` (Number) - total number of stacklinks
* `stack_ports_up` (Number) - total number of stack ports in up state
* `tool_ports_up` (Number) - total number of tool ports in up state

<a id="nestedatt--multi_upgrade_status--object_diff_report"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report`

Read-Only:

* `cards_up_differences` (Attributes List) - Diff in cards in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--cards_up_differences))
* `circuit_ports_up_differences` (Attributes List) - Diff in circuit ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--circuit_ports_up_differences))
* `gigasmart_ports_up_differences` (Attributes List) - Diff in gigasmart ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--gigasmart_ports_up_differences))
* `gsop_differences` (Attributes List) - Diff in gigasmart operations (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--gsop_differences))
* `hybrid_ports_up_differences` (Attributes List) - Diff in hybrid ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--hybrid_ports_up_differences))
* `inline_network_ports_up_differences` (Attributes List) - Diff in inlineNetwork ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--inline_network_ports_up_differences))
* `inline_tool_ports_up_differences` (Attributes List) - Diff in inlineTool ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--inline_tool_ports_up_differences))
* `maps_differences` (Attributes List) - Diff in maps (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--maps_differences))
* `network_ports_up_differences` (Attributes List) - Diff in network ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--network_ports_up_differences))
* `node_differences` (Attributes List) - Diff in nodes (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--node_differences))
* `stack_links_differences` (Attributes List) - Diff in stacklinks (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--stack_links_differences))
* `stack_ports_up_differences` (Attributes List) - Diff in stack ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--stack_ports_up_differences))
* `tool_ports_up_differences` (Attributes List) - Diff in tool ports in up state (see [below for nested schema](#nestedatt--multi_upgrade_status--object_diff_report--tool_ports_up_differences))

<a id="nestedatt--multi_upgrade_status--object_diff_report--cards_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.cards_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--circuit_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.circuit_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--gigasmart_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.gigasmart_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--gsop_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.gsop_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--hybrid_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.hybrid_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--inline_network_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.inline_network_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--inline_tool_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.inline_tool_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--maps_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.maps_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--network_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.network_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--node_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.node_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--stack_links_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.stack_links_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--stack_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.stack_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--object_diff_report--tool_ports_up_differences"></a>
### Nested Schema for `multi_upgrade_status.object_diff_report.tool_ports_up_differences`

Read-Only:

* `after_state` (String) - next state of the object
* `object_name` (String) - alias of the object
* `prev_state` (String) - previous state of the object

<a id="nestedatt--multi_upgrade_status--upgrade_summary"></a>
### Nested Schema for `multi_upgrade_status.upgrade_summary`

Read-Only:

* `num_initial_validation_complete` (Number) - total number of items in the initial validation complete state
* `num_install_complete` (Number) - total number of items in the installation complete state
* `num_nodes_success` (Number) - total number of items in the nodes success state
* `num_reload_complete` (Number) - total number of items in the reload complete state
* `num_uboot_complete` (Number) - total number of items in the Uboot complete state
* `num_upgrade_complete` (Number) - total number of items in the upgrade complete state
* `num_verification_complete` (Number) - total number of items in the verification complete state

