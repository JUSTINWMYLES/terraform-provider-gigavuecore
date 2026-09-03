---
page_title: "gigavuecore_get_expiring_count Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns number of FM Licenses that are expiring soon or are recently expired
---

# gigavuecore_get_expiring_count Data Source

Returns number of FM Licenses that are expiring soon or are recently expired

## Example Usage

```terraform
data "gigavuecore_get_expiring_count" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `count_` (String, computed)


