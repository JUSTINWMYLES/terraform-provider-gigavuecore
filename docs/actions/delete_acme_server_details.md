---
page_title: "gigavuecore_delete_acme_server_details Action - gigavuecore"
subcategory: ""
description: |-
  Delete ACME server details
---

# gigavuecore_delete_acme_server_details Action

Delete ACME server details

## Example Usage

```terraform
action "gigavuecore_delete_acme_server_details" "example" {
  config {
    alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the acme server


