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
    cluster_id = "example"
    pcap_delete_list = [{
      box_id    = "example"
      file_name = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `pcap_delete_list` (Attributes List, optional) (see [below for nested schema](#nestedatt--pcap_delete_list))

<a id="nestedatt--pcap_delete_list"></a>
### Nested Schema for `pcap_delete_list`

Required:

* `box_id` (String) - The chassis Box ID value
* `file_name` (String) - Name of the PCAP file

