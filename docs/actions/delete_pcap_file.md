---
page_title: "gigavuecore_delete_pcap_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete Pcap File
---

# gigavuecore_delete_pcap_file Action

Delete Pcap File

## Example Usage

```terraform
action "gigavuecore_delete_pcap_file" "example" {
  config {
    cluster_id       = "example"
    pcap_delete_list = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `pcap_delete_list` (List of Dynamic, optional)


