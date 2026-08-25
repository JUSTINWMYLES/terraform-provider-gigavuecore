---
page_title: "gigavuecore_add_session_aware_apf_field Action - gigavuecore"
subcategory: ""
description: |-
  Add a new session field to a SA-APF Profile
---

# gigavuecore_add_session_aware_apf_field Action

Add a new session field to a SA-APF Profile

## Example Usage

```terraform
action "gigavuecore_add_session_aware_apf_field" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    pos        = 1
    type       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target SA-APF Profile
* `cluster_id` (String, required) - Target Cluster ID
* `pos` (Number, optional) - Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported
* `type` (String, required)


