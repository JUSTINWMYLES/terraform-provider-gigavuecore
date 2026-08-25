---
page_title: "gigavuecore_delete_all_sip_whitelist_entries Action - gigavuecore"
subcategory: ""
description: |-
  Delete All SIP Whitelist Entries
---

# gigavuecore_delete_all_sip_whitelist_entries Action

Delete All SIP Whitelist Entries

## Example Usage

```terraform
action "gigavuecore_delete_all_sip_whitelist_entries" "example" {
  config {
    alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target SIP Whitelist


