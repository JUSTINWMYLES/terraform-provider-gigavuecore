---
page_title: "gigavuecore_tag List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all tags
---

# gigavuecore_tag List Resource

Load all tags

-> **Note:** This list resource requires Terraform 1.14 or later and is used through the `terraform query` command, not in configuration files.

## Example Usage

```terraform
list "gigavuecore_tag" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    page    = "example"
    sort    = "example"
    tag_key = "example"
    type    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `tag_key` (String, optional) - Filter based on provided tag keys
* `type` (String, optional) - If provided, returns all tags which matches the tag type


### Identity Attributes

The following identity attributes are exported for each matching result:

* `tag_key` (String, computed) - tag key


