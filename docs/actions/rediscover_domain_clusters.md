---
page_title: "gigavuecore_rediscover_domain_clusters Action - gigavuecore"
subcategory: ""
description: |-
  Rediscovers all managed clusters and standalone nodes under FM management
---

# gigavuecore_rediscover_domain_clusters Action

Rediscovers all managed clusters and standalone nodes under FM management

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_rediscover_domain_clusters" "example" {
  config {
    async      = true
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `async` (Boolean, optional) - if provided, the call returns immediately with the \[201 Accepted\] HTTP status code, and the discovery process will complete in the background
* `cluster_id` (String, optional) - if provided, only requested cluster is rediscovered. For Standalone devices, rediscovers that one node (identified by its nodeId). For clustered nodes, rediscovers entire cluster


