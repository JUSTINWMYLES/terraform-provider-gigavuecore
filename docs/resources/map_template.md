---
page_title: "gigavuecore_map_template Resource - gigavuecore"
subcategory: ""
description: |-
  Find map template by alias
---

# gigavuecore_map_template Resource

Find map template by alias

## Example Usage

```terraform
resource "gigavuecore_map_template" "example" {
  alias = null
  cluster_id = null
  comment = null
  roles = {}
  rules = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - unique map alias
* `cluster_id` (String, optional) - id of the defining cluster
* `comment` (String, optional)
* `roles` (Object({editors, owners, viewers}), optional) - Map Roles Container. Private class
  * `editors` (List(String), optional)
  * `owners` (List(String), optional)
  * `viewers` (List(String), optional)
* `rules` (Object({drop_rules, pass_rules}), required) - Map Rules Container. Private class
  * `drop_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)
  * `pass_rules` (Set(Object({bidi, comment, ip_rewrite, matches, rewrite, rule_id, vlan_tag})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed)
* `roles` (Object({editors, owners, viewers}), computed) - Map Roles Container. Private class
  * `editors` (List(String), optional)
  * `owners` (List(String), optional)
  * `viewers` (List(String), optional)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_map_template.example {alias}
```
