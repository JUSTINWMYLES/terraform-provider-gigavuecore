---
page_title: "gigavuecore_generate_gigasmart_dump Action - gigavuecore"
subcategory: ""
description: |-
  Generate Gigasmart Dump
---

# gigavuecore_generate_gigasmart_dump Action

Generate Gigasmart Dump

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_generate_gigasmart_dump" "example" {
  config {
    cluster_id = "example"
    engine_ids = ["example"]
    hostname   = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster Id for which dump is initiated
* `engine_ids` (List of String, required) - list of the Gigasmart engine ports for which dump is initiated
* `hostname` (String, required) - Hostname for which dump is initiated


