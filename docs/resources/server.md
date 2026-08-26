---
page_title: "gigavuecore_server Resource - gigavuecore"
subcategory: ""
description: |-
  Load ICAP Server by Alias
---

# gigavuecore_server Resource

Load ICAP Server by Alias

## Example Usage

```terraform
resource "gigavuecore_server" "example" {
  alias               = null
  comment             = null
  l3_address          = null
  l4_port             = null
  options_service_url = null
  reqmod_service_url  = null
  respmod_service_url = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Server Alias
* `comment` (String, optional) - Icap Server Comment
* `l3_address` (String, required) - Icap Server IP Address
* `l4_port` (Number, required) - Icap Server l4 Port
* `options_service_url` (String, optional) - Options Service URL
* `reqmod_service_url` (String, optional) - Request Modification Service URL
* `respmod_service_url` (String, optional) - Response Modification Service URL

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `comment` (String, computed) - Icap Server Comment
* `options_service_url` (String, computed) - Options Service URL
* `reqmod_service_url` (String, computed) - Request Modification Service URL
* `respmod_service_url` (String, computed) - Response Modification Service URL


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_server.example {alias}
```
