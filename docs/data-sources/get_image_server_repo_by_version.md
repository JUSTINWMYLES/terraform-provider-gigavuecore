---
page_title: "gigavuecore_get_image_server_repo_by_version Data Source - gigavuecore"
subcategory: ""
description: |-
  Find Images By Version
---

# gigavuecore_get_image_server_repo_by_version Data Source

Find Images By Version

## Example Usage

```terraform
data "gigavuecore_get_image_server_repo_by_version" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({images, version})), computed) - System Image File by version

