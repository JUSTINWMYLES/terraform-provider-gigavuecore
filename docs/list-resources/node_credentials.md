---
page_title: "gigavuecore_node_credentials List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Device Credentials
---

# gigavuecore_node_credentials List Resource

Load all Device Credentials

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

~> **Warning:** This list resource accepts attributes whose names indicate secrets (http_password), but list resource schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
list "gigavuecore_node_credentials" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    page = "example"
    sort = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)


### Identity Attributes

The following identity attributes are exported for each matching result:

* `device_address` (String, computed) - value of '0.0.0.0' represents the default (fallback) device credentials


