---
page_title: "gigavuecore_load_pcap_file Data Source - gigavuecore"
subcategory: ""
description: |-
  Download Pcap File
---

# gigavuecore_load_pcap_file Data Source

Download Pcap File

## Example Usage

```terraform
data "gigavuecore_load_pcap_file" "example" {
  box_id     = "example"
  cluster_id = "example"
  filename   = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID range from 1 to 64(inclusive)
* `cluster_id` (String, required) - Target Cluster ID
* `filename` (String, required) - pcap filename to download


