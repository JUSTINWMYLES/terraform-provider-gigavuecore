---
page_title: "gigavuecore_get_ping_result Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Ping Result
---

# gigavuecore_get_ping_result Data Source

Get Ping Result

## Example Usage

```terraform
data "gigavuecore_get_ping_result" "example" {
  cluster_id = "example"
  eport      = "example"
  ip         = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `eport` (String, required) - GigaSMART engine port
* `ip` (String, required) - ipv4 or hostname

### Attributes

In addition to all arguments above, the following attributes are exported:

* `ping_response` (String, computed)


