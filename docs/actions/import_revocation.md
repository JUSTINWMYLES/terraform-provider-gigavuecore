---
page_title: "gigavuecore_import_revocation Action - gigavuecore"
subcategory: ""
description: |-
  Import FM License Revocation
---

# gigavuecore_import_revocation Action

Import FM License Revocation

## Example Usage

```terraform
action "gigavuecore_import_revocation" "example" {
  config {
    revocation_file_name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `revocation_file_name` (String, required) - User uploaded license file for revocation
