---
page_title: "gigavuecore_validate_updated_policy Action - gigavuecore"
subcategory: ""
description: |-
  Validate Policy
---

# gigavuecore_validate_updated_policy Action

Validate Policy

## Example Usage

```terraform
action "gigavuecore_validate_updated_policy" "example" {
  config {
    body_name      = "example"
    comment        = "example"
    dest_ports     = null
    gigasmart_info = null
    name           = "example"
    policy_id      = "example"
    priority       = "example"
    rule_criteria  = null
    rules_info     = null
    src_ports      = null
    tags           = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_name` (String, optional)
* `comment` (String, optional)
* `dest_ports` (List of Dynamic, optional)
* `gigasmart_info` (List of Dynamic, optional)
* `name` (String, optional) - policy name
* `policy_id` (String, optional)
* `priority` (String, optional)
* `rule_criteria` (List of Dynamic, optional)
* `rules_info` (List of Dynamic, optional)
* `src_ports` (List of Dynamic, optional)
* `tags` (List of Dynamic, optional)


