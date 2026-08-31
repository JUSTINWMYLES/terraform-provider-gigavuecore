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
  box_id     = "example"
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID range from 1 to 64(inclusive). all is applicable
* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) - List of available pcap files (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `box_id` (String) - The chassis Box ID value
* `capture_files` (Attributes List) - List of available pcap files (see [below for nested schema](#nestedatt--items--capture_files))

<a id="nestedatt--items--capture_files"></a>
### Nested Schema for `items.capture_files`

Read-Only:

* `filename` (String) - Name of the PCAP file
* `size` (Number) - Size of file in bytes
* `timestamp` (String) - File creation date and time in ISO 8601 format

