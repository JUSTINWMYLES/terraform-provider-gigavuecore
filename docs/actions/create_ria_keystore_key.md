---
page_title: "gigavuecore_create_ria_keystore_key Action - gigavuecore"
subcategory: ""
description: |-
  add key and/or certificate to the keystore for two devices needed for Resilient Inline SSL
---

# gigavuecore_create_ria_keystore_key Action

add key and/or certificate to the keystore for two devices needed for Resilient Inline SSL

## Example Usage

```terraform
action "gigavuecore_create_ria_keystore_key" "example" {
  config {
    alias = "example"
    certificate = null
    cluster_names = [ "example" ]
    comment = "example"
    key = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `certificate` (Dynamic, optional)
* `cluster_names` (List(String), optional)
* `comment` (String, optional)
* `key` (Dynamic, optional)
