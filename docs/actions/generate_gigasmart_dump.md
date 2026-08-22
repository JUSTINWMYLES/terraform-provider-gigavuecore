---
page_title: "gigavuecore_generate_gigasmart_dump Action - gigavuecore"
subcategory: ""
description: |-
  Generate Gigasmart Dump
---

# gigavuecore_generate_gigasmart_dump Action

Generate Gigasmart Dump

## Example Usage

```terraform
action "gigavuecore_generate_gigasmart_dump" "example" {
  config {
    cluster_id = "example"
    engine_ids = [ "example" ]
    hostname = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster Id for which dump is initiated
* `engine_ids` (List(String), required) - list of the Gigasmart engine ports for which dump is initiated
* `hostname` (String, required) - Hostname for which dump is initiated
