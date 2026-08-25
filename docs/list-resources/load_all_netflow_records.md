---
page_title: "gigavuecore_load_all_netflow_records List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all defined Netflow Templates
---

# gigavuecore_load_all_netflow_records List Resource

Load all defined Netflow Templates

## Example Usage

```terraform
list "gigavuecore_load_all_netflow_records" "example" {
  provider = gigavuecore
  limit    = 100
}

```
