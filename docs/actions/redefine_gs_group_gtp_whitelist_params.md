---
page_title: "gigavuecore_redefine_gs_group_gtp_whitelist_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's GTP Whitelist Params
---

# gigavuecore_redefine_gs_group_gtp_whitelist_params Action

Redefine GS Group's GTP Whitelist Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gtp_whitelist_params" "example" {
  config {
    alias            = "example"
    cluster_id       = "example"
    multi_whitelists = [ "example" ]
    whitelist        = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `multi_whitelists` (List of String, optional) - Alias/Aliases of referenced GTP Whitelists.
* `whitelist` (String, required) - Alias of referenced GTP Whitelist.Deprecated since H 5.12


