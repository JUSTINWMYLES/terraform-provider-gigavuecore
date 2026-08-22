---
page_title: "gigavuecore_load_src_port_meta_data Action - gigavuecore"
subcategory: ""
description: |-
  Load source port metadata for traffic flows
---

# gigavuecore_load_src_port_meta_data Action

Load source port metadata for traffic flows

## Example Usage

```terraform
action "gigavuecore_load_src_port_meta_data" "example" {
  config {
    source_details = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `source_details` (List(Dynamic), optional)
