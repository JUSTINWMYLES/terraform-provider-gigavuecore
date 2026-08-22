---
page_title: "gigavuecore_delete_report Action - gigavuecore"
subcategory: ""
description: |-
  Delete CDA Audit Report
---

# gigavuecore_delete_report Action

Delete CDA Audit Report

## Example Usage

```terraform
action "gigavuecore_delete_report" "example" {
  config {
    file_name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `file_name` (String, required) - Name of the CDA Audit Report
