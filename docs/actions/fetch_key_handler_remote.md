---
page_title: "gigavuecore_fetch_key_handler_remote Action - gigavuecore"
subcategory: ""
description: |-
  Fetch Key Handler from Remote
---

# gigavuecore_fetch_key_handler_remote Action

Fetch Key Handler from Remote

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_fetch_key_handler_remote" "example" {
  config {
    alias    = "example"
    hostname = "example"
    password = "example"
    path     = "example"
    protocol = "scp"
    username = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `hostname` (String, required) - server address
* `password` (String, optional) - password to use for server login
* `path` (String, required) - file path on server
* `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
* `username` (String, optional) - user name to use for server login


