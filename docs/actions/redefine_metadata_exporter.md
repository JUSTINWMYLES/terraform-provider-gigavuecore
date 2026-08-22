---
page_title: "gigavuecore_redefine_metadata_exporter Action - gigavuecore"
subcategory: ""
description: |-
  Redefine a metadata exporter
---

# gigavuecore_redefine_metadata_exporter Action

Redefine a metadata exporter

## Example Usage

```terraform
action "gigavuecore_redefine_metadata_exporter" "example" {
  config {
    alias = "example"
    application_profiles = [ "example" ]
    body_alias = "example"
    cef = null
    description = "example"
    destination = null
    max_pkt_size = 1
    mobility_sam = null
    monitor = null
    netflow = null
    snmp = null
    source = null
    type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - metadata exporter alias
* `application_profiles` (List(String), optional) - application profile aliases to attach to the exporter
* `body_alias` (String, required)
* `cef` (Dynamic, optional)
* `description` (String, optional)
* `destination` (Dynamic, optional)
* `max_pkt_size` (Number, optional)
* `mobility_sam` (Dynamic, optional)
* `monitor` (Dynamic, optional)
* `netflow` (Dynamic, optional)
* `snmp` (Dynamic, optional)
* `source` (Dynamic, optional)
* `type` (String, optional)
