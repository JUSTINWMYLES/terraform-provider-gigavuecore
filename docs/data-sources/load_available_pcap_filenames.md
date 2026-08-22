---
page_title: "gigavuecore_load_available_pcap_filenames Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Available PCAP Filenames
---

# gigavuecore_load_available_pcap_filenames Data Source

Load Available PCAP Filenames

## Example Usage

```terraform
data "gigavuecore_load_available_pcap_filenames" "example" {
  box_id = null
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID range from 1 to 64(inclusive). all is applicable
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({box_id, capture_files})), computed) - List of available pcap files

