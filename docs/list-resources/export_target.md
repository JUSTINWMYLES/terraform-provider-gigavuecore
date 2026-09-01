---
page_title: "gigavuecore_export_target List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all External Export Server
---

# gigavuecore_export_target List Resource

Load all External Export Server

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

~> **Warning:** This list resource accepts attributes whose names indicate secrets (password), but list resource schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
list "gigavuecore_export_target" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `export_target_alias` (String, computed) - Alias for external export target server


