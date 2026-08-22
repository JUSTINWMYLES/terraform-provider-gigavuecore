---
page_title: "gigavuecore_get_policy_rules Action - gigavuecore"
subcategory: ""
description: |-
  Retrieve policy rules for a given alias
---

# gigavuecore_get_policy_rules Action

Retrieve policy rules for a given alias

## Example Usage

```terraform
action "gigavuecore_get_policy_rules" "example" {
  config {
    alias = "example"
    body = "example"
    flow_alias = "example"
    is_draft_version = true
    page = "example"
    sort = "example"
    src_alias = "example"
    sub_flow_alias = "example"
    text_search = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Policy alias
* `body` (Dynamic, required) - Rule Filter Info
* `flow_alias` (String, optional) - Flow alias
* `is_draft_version` (Bool, optional) - Load rules from the deployed draft policy
* `page` (String, optional) - Page request string
* `sort` (String, optional) - Sort order
* `src_alias` (String, optional) - Source alias
* `sub_flow_alias` (String, optional) - Sub-Flow alias
* `text_search` (String, optional) - Text search filter
