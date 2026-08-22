---
page_title: "gigavuecore_import_manual_topology_entity Action - gigavuecore"
subcategory: ""
description: |-
  Imports topology nodes and links from the information provided in csv file. For importing network devices "Tool Name", "Vendor", "Type", "Model", "Comment" and for tools the additional properties like "Max Throughput(Gbps)","Storage Capacity(TB)", "Compression Ratio" column headers are required. Similarly for links "Source Device"(device alias/hostname), "Source Cluster Name", "Source Device Type"(gigamon/manual), "Source EndPoint Type"(Port/Gigastream), "Source EndPoint Alias"(port/gigastream alias), "Destination Device", "Destination Cluster Name", "Destination Device Type", "Destination EndPoint Alias", "Destination EndPoint Type", "Link Type"(port/gigaStream) are required in the csv file.
---

# gigavuecore_import_manual_topology_entity Action

Imports topology nodes and links from the information provided in csv file. For importing network devices "Tool Name", "Vendor", "Type", "Model", "Comment" and for tools the additional properties like "Max Throughput(Gbps)","Storage Capacity(TB)", "Compression Ratio" column headers are required. Similarly for links "Source Device"(device alias/hostname), "Source Cluster Name", "Source Device Type"(gigamon/manual), "Source EndPoint Type"(Port/Gigastream), "Source EndPoint Alias"(port/gigastream alias), "Destination Device", "Destination Cluster Name", "Destination Device Type", "Destination EndPoint Alias", "Destination EndPoint Type", "Link Type"(port/gigaStream) are required in the csv file.

## Example Usage

```terraform
action "gigavuecore_import_manual_topology_entity" "example" {
  config {
    input = "example"
    type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `input` (String, required) - User uploaded file, only csv format is supported
* `type` (String, required) - Topology Entity Type
