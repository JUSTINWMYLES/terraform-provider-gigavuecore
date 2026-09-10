---
page_title: "gigavuecore_config Resource - gigavuecore"
subcategory: ""
description: |-
  Upload bulk replicate config file
---

# gigavuecore_config Resource

Upload bulk replicate config file

## Example Usage

```terraform
resource "gigavuecore_config" "example" {
  cluster_ids = ["example"]
  filename    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_ids` (List of String, optional) - Cluster Ids the config file needs to be applied
* `filename` (String, required) - Bulk Replicate config file name

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed) - User comments for the config file
* `created_ts` (Number, computed) - File creation timestamp in UTC milliseconds
* `created_ts_utc` (String, computed) - File creation timestamp in ISO 8601 format
* `family` (String, computed) - \['H' or 'G'\]: identifies whether this is an H-series or a G-series Config file

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
terraform import gigavuecore_config.example {filename}
```
