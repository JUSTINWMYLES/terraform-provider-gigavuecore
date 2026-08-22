---
page_title: "gigavuecore_import_policy Action - gigavuecore"
subcategory: ""
description: |-
  Import Policy
---

# gigavuecore_import_policy Action

Import Policy

## Example Usage

```terraform
action "gigavuecore_import_policy" "example" {
  config {
    criteria_bindings = null
    name = "example"
    packet_transformation_bindings = null
    priority = true
    rules = null
    source_bindings = null
    sources = null
    tags = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `criteria_bindings` (Dynamic, optional)
* `name` (String, optional)
* `packet_transformation_bindings` (Dynamic, optional)
* `priority` (Bool, optional)
* `rules` (List(Dynamic), optional)
* `source_bindings` (Dynamic, optional)
* `sources` (List(Dynamic), optional)
* `tags` (List(Dynamic), optional)
