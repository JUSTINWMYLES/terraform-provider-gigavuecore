---
page_title: "gigavuecore_hsm_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load HSM Group by alias
---

# gigavuecore_hsm_group Resource

Load HSM Group by alias

## Example Usage

```terraform
resource "gigavuecore_hsm_group" "example" {
  alias              = "example"
  comment            = "example"
  hsms               = [ "example" ]
  operational_status = "example"
  type               = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - hsm group alias
* `comment` (String, optional) - Hsm Group Comment
* `hsms` (List of String, optional) - alias of hsm in hsm group
* `operational_status` (String, optional) - operational status of hsm group
* `type` (String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `comment` (String, computed) - Hsm Group Comment
* `hsms` (List of String, computed) - alias of hsm in hsm group
* `operational_status` (String, computed) - operational status of hsm group


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_hsm_group.example {alias}
```
