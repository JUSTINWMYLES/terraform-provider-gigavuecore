---
page_title: "gigavuecore_delete_ntp_auth_key Action - gigavuecore"
subcategory: ""
description: |-
  Delete NTP Auth Key by alias
---

# gigavuecore_delete_ntp_auth_key Action

Delete NTP Auth Key by alias

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_ntp_auth_key" "example" {
  config {
    auth_key_id = 0
    cluster_id  = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `auth_key_id` (Number, required) - NTP Authentication Key Id
* `cluster_id` (String, required) - Target Cluster ID


