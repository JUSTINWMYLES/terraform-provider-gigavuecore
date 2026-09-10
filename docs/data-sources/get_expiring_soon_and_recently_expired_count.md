---
page_title: "gigavuecore_get_expiring_soon_and_recently_expired_count Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns number of floating licenses that are expiring soon or are recently expired
---

# gigavuecore_get_expiring_soon_and_recently_expired_count Data Source

Returns number of floating licenses that are expiring soon or are recently expired

## Example Usage

```terraform
data "gigavuecore_get_expiring_soon_and_recently_expired_count" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `expiring_grace_activations` (String, computed)
* `expiring_grace_licenses` (String, computed)
* `expiring_soon_activations` (String, computed)
* `expiring_soon_licenses` (String, computed)


