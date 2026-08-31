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
  alias               = "example"
  cluster_id          = "example"
  comment             = "example"
  l3_address          = "example"
  l4_port             = 0
  options_service_url = "example"
  reqmod_service_url  = "example"
  respmod_service_url = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Server Alias
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional) - Icap Server Comment
* `l3_address` (String, required) - Icap Server IP Address
* `l4_port` (Number, required) - Icap Server l4 Port
* `options_service_url` (String, optional) - Options Service URL
* `reqmod_service_url` (String, optional) - Request Modification Service URL
* `respmod_service_url` (String, optional) - Response Modification Service URL

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_server.example {alias}/{cluster_id}
```
