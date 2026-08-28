---
page_title: "gigavuecore_fetch_key_map Action - gigavuecore"
subcategory: ""
description: |-
  Fetch Key Maps from Remote
---

# gigavuecore_fetch_key_map Action

Fetch Key Maps from Remote

## Example Usage

```terraform
action "gigavuecore_fetch_key_map" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    hostname   = "example"
    password   = "example"
    path       = "example"
    protocol   = "example"
    username   = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
* `hostname` (String, required) - server address
* `password` (String, optional) - password to use for server login
* `path` (String, required) - file path on server
* `protocol` (String, required) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
* `username` (String, optional) - user name to use for server login


