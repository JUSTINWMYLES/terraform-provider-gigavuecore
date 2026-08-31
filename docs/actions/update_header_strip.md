---
page_title: "gigavuecore_update_header_strip Action - gigavuecore"
subcategory: ""
description: |-
  Update header strip for target box
---

# gigavuecore_update_header_strip Action

Update header strip for target box

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_header_strip" "example" {
  config {
    add_labels    = ["example"]
    body_box_id   = "example"
    box_id        = "example"
    cluster_id    = "example"
    delete_labels = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `add_labels` (List of String, optional) - mpls ids, valid and required. Range can be specified. Example:1..200
* `body_box_id` (String, required) - device box id. valid range 1 - 64.
* `box_id` (String, required) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target cluster ID.
* `delete_labels` (List of String, optional) - mpls ids, valid and required. Range can be specified. Example:1..200


