---
page_title: "gigavuecore_server Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new ICAP Server
---

# gigavuecore_server Resource

Create a new ICAP Server

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_server.example {alias}/{cluster_id}
```
