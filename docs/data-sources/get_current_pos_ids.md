---
page_title: "gigavuecore_get_current_pos_ids Data Source - gigavuecore"
subcategory: ""
description: |-
  Gives the set of POS IDs associated with currently unexpired (including recently expired) volume licenses
---

# gigavuecore_get_current_pos_ids Data Source

Gives the set of POS IDs associated with currently unexpired (including recently expired) volume licenses

## Example Usage

```terraform
data "gigavuecore_get_current_pos_ids" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List of String, computed)


