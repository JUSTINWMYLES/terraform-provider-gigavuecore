---
page_title: "gigavuecore_port_filter Resource - gigavuecore"
subcategory: ""
description: |-
  get filter of a Port by port ID
---

# gigavuecore_port_filter Resource

get filter of a Port by port ID

## Example Usage

```terraform
resource "gigavuecore_port_filter" "example" {
  port = null
  rules = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `port` (String, required)
* `rules` (Object({drop_rules, pass_rules}), required) - Port Filter Rules Container. Private class
  * `drop_rules` (Set(Object({comment, matches, rule_id})), optional)
  * `pass_rules` (Set(Object({comment, matches, rule_id})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `port_id` (String, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_filter.example {port_id}
```
