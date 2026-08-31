---
page_title: "gigavuecore_download_tag_resources Action - gigavuecore"
subcategory: ""
description: |-
  download the tag resources as a file
---

# gigavuecore_download_tag_resources Action

download the tag resources as a file

## Example Usage

```terraform
action "gigavuecore_download_tag_resources" "example" {
  config {
    cluster_names      = [ "example" ]
    file_name          = "example"
    tag_present        = true
    tag_resource_types = [ "Cluster" ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_names` (List of String, optional) - list of ClusterIds to download the resources, if not specified all the supported resources from the clusters would be downloaded
* `file_name` (String, required) - specifies the file name
* `tag_present` (Boolean, optional) - if enabled only the resources with tags would be present in the downloaded file
* `tag_resource_types` (List of String, optional) - list of TagResourceType to filter by, if not specified all the supported resources (port, map, cluster) would be downloaded


