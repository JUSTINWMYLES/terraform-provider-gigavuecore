---
page_title: "gigavuecore_generate_sysdump Action - gigavuecore"
subcategory: ""
description: |-
  Generate New Sysdump File
---

# gigavuecore_generate_sysdump Action

Generate New Sysdump File

## Example Usage

```terraform
action "gigavuecore_generate_sysdump" "example" {
  config {
    box_id = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID range from 1 to 64(inclusive). all is applicable
* `cluster_id` (String, required) - Target Cluster ID
