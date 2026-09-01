---
page_title: "gigavuecore_hsm List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all HSM
---

# gigavuecore_hsm List Resource

Load all HSM

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

~> **Warning:** This list resource accepts attributes whose names indicate secrets (partition_password, server_password), but list resource schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
list "gigavuecore_hsm" "example" {
  provider = gigavuecore
  limit    = 100
}
```
## Schema

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed) - Hsm alias


