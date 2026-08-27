---
page_title: "gigavuecore_load_all_image_servers Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Image Servers
---

# gigavuecore_load_all_image_servers Data Source

Load all Image Servers

## Example Usage

```terraform
data "gigavuecore_load_all_image_servers" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `address` (String)
* `alias` (String) - unique alias for this image server
* `type` (String)
* `user_pwd` (String) - user password to use for server login. not applicable for 'tftp'
* `username` (String) - username to use for server login. not applicable for 'tftp'

