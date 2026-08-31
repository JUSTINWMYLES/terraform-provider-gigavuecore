---
page_title: "gigavuecore_upload_certificate_ssl_node_certificate_upload_name Action - gigavuecore"
subcategory: ""
description: |-
  upload a new certificate and add it to the device' CA List the device
---

# gigavuecore_upload_certificate_ssl_node_certificate_upload_name Action

upload a new certificate and add it to the device' CA List the device

## Example Usage

```terraform
action "gigavuecore_upload_certificate_ssl_node_certificate_upload_name" "example" {
  config {
    add_to_ca_list = true
    certificate    = "example"
    cluster_id     = "example"
    comment        = "example"
    name           = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `add_to_ca_list` (Boolean, optional) - set this true to directly push the certificate to device's CA List
* `certificate` (String, required) - User uploaded file, only .crt/.cert or .pem format is supported
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional) - a short description of the certificate
* `name` (String, required) - the certificate name


