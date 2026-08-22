---
page_title: "gigavuecore_list_env Data Source - gigavuecore"
subcategory: ""
description: |-
  List all unified resource environments
---

# gigavuecore_list_env Data Source

List all unified resource environments

## Example Usage

```terraform
data "gigavuecore_list_env" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({env})), computed)

