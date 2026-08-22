---
page_title: "gigavuecore_upload_acme_server_config_details Action - gigavuecore"
subcategory: ""
description: |-
  upload ACME server details
---

# gigavuecore_upload_acme_server_config_details Action

upload ACME server details

## Example Usage

```terraform
action "gigavuecore_upload_acme_server_config_details" "example" {
  config {
    acme_server_url = "example"
    alias = "example"
    certificate = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `acme_server_url` (String, required) - ACME server url
* `alias` (String, required) - ACME server alias
* `certificate` (String, required) - User uploaded file, only .crt/.cert or .pem format is supported
