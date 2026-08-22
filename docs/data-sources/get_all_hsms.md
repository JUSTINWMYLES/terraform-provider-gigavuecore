---
page_title: "gigavuecore_get_all_hsms Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all HSM
---

# gigavuecore_get_all_hsms Data Source

Load all HSM

## Example Usage

```terraform
data "gigavuecore_get_all_hsms" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, esn, hsm_ip, hsm_port, kneti, operational_status, partition_label, partition_password, server_password, server_username, type})), computed)

