---
page_title: "gigavuecore_cluster_config_edit_leader_preference Action - gigavuecore"
subcategory: ""
description: |-
  update the leader preference of the specified member
---

# gigavuecore_cluster_config_edit_leader_preference Action

update the leader preference of the specified member

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_cluster_config_edit_leader_preference" "example" {
  config {
    box_id            = "example"
    cluster_id        = "example"
    leader_preference = 0
    master_preference = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - box id of the device whose leader preference needs to be changed
* `cluster_id` (String, required) - cluster id to which the device belongs
* `leader_preference` (Number, optional) - leader election preference rank, 1 to 9 to exclude leadership
* `master_preference` (Number, optional) - master election preference rank, 1 to 9 to exclude mastership. (deprecated: use leaderPreference)


