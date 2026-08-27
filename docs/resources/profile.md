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
  alias               = null
  cluster_id          = null
  exceed_action       = null
  http_req_buf        = null
  inactivity_timeout  = null
  preview             = null
  resp_mod            = null
  resp_timeout        = null
  resp_timeout_action = null
  server_group        = null
  src_max_l4_port     = null
  src_min_l4_port     = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `exceed_action` (String, computed) - Icap Profile action incase of Http request buffer exceeded
* `inactivity_timeout` (Number, computed) - Icap Inactivity timeout in minutes
* `preview` (Number, computed) - Icap Preview bytes in KB
* `resp_mod` (String, computed) - Icap Response Modification Enable\|Disable
* `resp_timeout` (Number, computed) - Icap Server Response Timeout value in seconds
* `resp_timeout_action` (String, computed) - Response Timeout action Drop\|Bypass


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_profile.example {alias}
```
