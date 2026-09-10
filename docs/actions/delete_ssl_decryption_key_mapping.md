---
page_title: "gigavuecore_delete_ssl_decryption_key_mapping Action - gigavuecore"
subcategory: ""
description: |-
  Delete an Endpoint-to-Key mapping
---

# gigavuecore_delete_ssl_decryption_key_mapping Action

Delete an Endpoint-to-Key mapping

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_ssl_decryption_key_mapping" "example" {
  config {
    alias          = "example"
    cluster_id     = "example"
    endpoint_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target KeyMap
* `cluster_id` (String, required) - Target Cluster ID
* `endpoint_alias` (String, required) - alias of the target SSL Decryption Endpoint


