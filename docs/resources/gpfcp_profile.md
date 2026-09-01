---
page_title: "gigavuecore_gpfcp_profile Resource - gigavuecore"
subcategory: ""
description: |-
  new in H 6.8
---

# gigavuecore_gpfcp_profile Resource

new in H 6.8

## Example Usage

```terraform
resource "gigavuecore_gpfcp_profile" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  g_profiles = [{
    comment = "example"
    g_interface = {
      ip_addresses = ["example"]
    }
    ip_interface = "example"
    node_type    = "control"
    port_list    = [0]
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the Gpfcp Profile
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional) - Description of the Gpfcp Profile
* `g_profiles` (Attributes List, required) (see [below for nested schema](#nestedatt--g_profiles))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--g_profiles"></a>
### Nested Schema for `g_profiles`

Required:

* `g_interface` (Attributes) (see [below for nested schema](#nestedatt--g_profiles--g_interface))
* `ip_interface` (String)
* `node_type` (String)
* `port_list` (List of Number)

Optional:

* `comment` (String) - Description of the Gpfcp Profile Rule

<a id="nestedatt--g_profiles--g_interface"></a>
### Nested Schema for `g_profiles.g_interface`

Optional:

* `ip_addresses` (List of String)
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
terraform import gigavuecore_gpfcp_profile.example {alias}/{cluster_id}
```
