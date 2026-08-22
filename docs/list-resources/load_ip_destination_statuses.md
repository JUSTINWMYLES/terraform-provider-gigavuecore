---
page_title: "gigavuecore_load_ip_destination_statuses List Resource - gigavuecore"
subcategory: ""
description: |-
  Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup
---

# gigavuecore_load_ip_destination_statuses List Resource

Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup

## Example Usage

```terraform
list "gigavuecore_load_ip_destination_statuses" "example" {
  provider = gigavuecore
  limit = 100
}

```

## Schema

### Arguments

The following arguments are supported:

