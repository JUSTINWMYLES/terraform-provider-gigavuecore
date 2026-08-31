---
page_title: "gigavuecore_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Load ICAP Profiles by Alias
---

# gigavuecore_profile Resource

Load ICAP Profiles by Alias

## Example Usage

```terraform
resource "gigavuecore_profile" "example" {
  alias               = "example"
  cluster_id          = "example"
  exceed_action       = "drop"
  http_req_buf        = 10
  inactivity_timeout  = 2
  preview             = 0
  resp_mod            = "enable"
  resp_timeout        = 5
  resp_timeout_action = "drop"
  server_group        = "example"
  src_max_l4_port     = 10000
  src_min_l4_port     = 10000
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Icap Alias
* `cluster_id` (String, required) - id of the defining cluster
* `exceed_action` (String, optional) - Icap Profile action incase of Http request buffer exceeded
* `http_req_buf` (Number, required) - Icap Profile Http request buffer in KB
* `inactivity_timeout` (Number, optional) - Icap Inactivity timeout in minutes
* `preview` (Number, optional) - Icap Preview bytes in KB
* `resp_mod` (String, optional) - Icap Response Modification Enable\|Disable
* `resp_timeout` (Number, optional) - Icap Server Response Timeout value in seconds
* `resp_timeout_action` (String, optional) - Response Timeout action Drop\|Bypass
* `server_group` (String, required) - Icap Server Group Alias
* `src_max_l4_port` (Number, required) - Icap Service Source l4 port maximum
* `src_min_l4_port` (Number, required) - Icap Service Source l4 port minimum

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_profile.example {alias}/{cluster_id}
```
