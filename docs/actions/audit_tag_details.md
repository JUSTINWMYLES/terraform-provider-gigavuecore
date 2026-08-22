---
page_title: "gigavuecore_audit_tag_details Action - gigavuecore"
subcategory: ""
description: |-
  Reapply cluster and node level tags to ports
---

# gigavuecore_audit_tag_details Action

Reapply cluster and node level tags to ports

## Example Usage

```terraform
action "gigavuecore_audit_tag_details" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
