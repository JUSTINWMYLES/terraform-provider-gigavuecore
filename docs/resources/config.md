---
page_title: "gigavuecore_config Resource - gigavuecore"
subcategory: ""
description: |-
  Find bulk replicate config file by file name
---

# gigavuecore_config Resource

Find bulk replicate config file by file name

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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_config.example {filename}
```
