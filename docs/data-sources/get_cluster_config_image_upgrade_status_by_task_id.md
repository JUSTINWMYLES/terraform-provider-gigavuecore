---
page_title: "gigavuecore_get_cluster_config_image_upgrade_status_by_task_id Data Source - gigavuecore"
subcategory: ""
description: |-
  get the cluster configuration imageUpgrade status by taskId
---

# gigavuecore_get_cluster_config_image_upgrade_status_by_task_id Data Source

get the cluster configuration imageUpgrade status by taskId

## Example Usage

```terraform
data "gigavuecore_get_cluster_config_image_upgrade_status_by_task_id" "example" {
  task_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `task_id` (String, required) - taskId of the task

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_id` (String, computed) - Id of the Target cluster
* `cluster_log` (List(Object({cluster_id, message, node_id, time})), computed) - list of the cluster Log
* `end_time` (String, computed) - End Time in UTC format
* `node_upgrade_status` (List(Object({current_step, current_version, file_name, file_server, hostname, node_id, node_log, num_steps_in_task, prior_version, status, step_status})), computed) - list of the nodeUpgrade status
* `object_counts` (Object({cards_up, circuit_ports_up, gigasmart_ports_up, gsops, hybrid_ports_up, inline_network_ports_up, inline_tool_ports_up, maps, network_ports_up, nodes, stack_links, stack_ports_up, tool_ports_up}), computed) - Cluster objectCounts
  * `cards_up` (Number, computed) - total number of cards in up state
  * `circuit_ports_up` (Number, computed) - total number of circuit ports in up state
  * `gigasmart_ports_up` (Number, computed) - total number of gigasmart ports in up state
  * `gsops` (Number, computed) - total number of gigasmart operations
  * `hybrid_ports_up` (Number, computed) - total number of hybrid ports in up state
  * `inline_network_ports_up` (Number, computed) - total number of inlineNetwork ports in up state
  * `inline_tool_ports_up` (Number, computed) - total number of inlineTool ports in up state
  * `maps` (Number, computed) - total number of maps
  * `network_ports_up` (Number, computed) - total number of network ports in up state
  * `nodes` (Number, computed) - total number of nodes
  * `stack_links` (Number, computed) - total number of stacklinks
  * `stack_ports_up` (Number, computed) - total number of stack ports in up state
  * `tool_ports_up` (Number, computed) - total number of tool ports in up state
* `object_counts_expected` (Object({cards_up, circuit_ports_up, gigasmart_ports_up, gsops, hybrid_ports_up, inline_network_ports_up, inline_tool_ports_up, maps, network_ports_up, nodes, stack_links, stack_ports_up, tool_ports_up}), computed) - Cluster objectCounts Expected
  * `cards_up` (Number, computed) - total number of cards in up state
  * `circuit_ports_up` (Number, computed) - total number of circuit ports in up state
  * `gigasmart_ports_up` (Number, computed) - total number of gigasmart ports in up state
  * `gsops` (Number, computed) - total number of gigasmart operations
  * `hybrid_ports_up` (Number, computed) - total number of hybrid ports in up state
  * `inline_network_ports_up` (Number, computed) - total number of inlineNetwork ports in up state
  * `inline_tool_ports_up` (Number, computed) - total number of inlineTool ports in up state
  * `maps` (Number, computed) - total number of maps
  * `network_ports_up` (Number, computed) - total number of network ports in up state
  * `nodes` (Number, computed) - total number of nodes
  * `stack_links` (Number, computed) - total number of stacklinks
  * `stack_ports_up` (Number, computed) - total number of stack ports in up state
  * `tool_ports_up` (Number, computed) - total number of tool ports in up state
* `object_diff_report` (Object({cards_up_differences, circuit_ports_up_differences, gigasmart_ports_up_differences, gsop_differences, hybrid_ports_up_differences, inline_network_ports_up_differences, inline_tool_ports_up_differences, maps_differences, network_ports_up_differences, node_differences, stack_links_differences, stack_ports_up_differences, tool_ports_up_differences}), computed) - Cluster object Diff report
  * `cards_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in cards in up state
  * `circuit_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in circuit ports in up state
  * `gigasmart_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in gigasmart ports in up state
  * `gsop_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in gigasmart operations
  * `hybrid_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in hybrid ports in up state
  * `inline_network_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in inlineNetwork ports in up state
  * `inline_tool_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in inlineTool ports in up state
  * `maps_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in maps
  * `network_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in network ports in up state
  * `node_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in nodes
  * `stack_links_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in stacklinks
  * `stack_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in stack ports in up state
  * `tool_ports_up_differences` (List(Object({after_state, object_name, prev_state})), computed) - Diff in tool ports in up state
* `start_time` (String, computed) - Start Time in UTC format
* `status` (String, computed) - Cluster Image Upgrade Status
* `task_name` (String, computed) - name of the task
* `upgrade_summary` (Object({num_initial_validation_complete, num_install_complete, num_nodes_success, num_reload_complete, num_uboot_complete, num_upgrade_complete, num_verification_complete}), computed) - Cluster upgrade Summary
  * `num_initial_validation_complete` (Number, computed) - total number of items in the initial validation complete state
  * `num_install_complete` (Number, computed) - total number of items in the installation complete state
  * `num_nodes_success` (Number, computed) - total number of items in the nodes success state
  * `num_reload_complete` (Number, computed) - total number of items in the reload complete state
  * `num_uboot_complete` (Number, computed) - total number of items in the Uboot complete state
  * `num_upgrade_complete` (Number, computed) - total number of items in the upgrade complete state
  * `num_verification_complete` (Number, computed) - total number of items in the verification complete state

