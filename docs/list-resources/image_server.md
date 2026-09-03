---
page_title: "gigavuecore_image_server List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Image Servers
---

# gigavuecore_image_server List Resource

Load all Image Servers

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_image_server" "example" {
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

* `alias` (String, computed) - unique alias for this image server


