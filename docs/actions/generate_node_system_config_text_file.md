---
page_title: "gigavuecore_generate_node_system_config_text_file Action - gigavuecore"
subcategory: ""
description: |-
  Generate active running text configuration file
---

# gigavuecore_generate_node_system_config_text_file Action

Generate active running text configuration file

## Example Usage

```terraform
action "gigavuecore_generate_node_system_config_text_file" "example" {
  config {
    cluster_id = "example"
    destination = {
      hostname = "example"
      password = "example"
      path     = "example"
      username = "example"
    }
    filename     = "example"
    only_traffic = true
    protocol     = "example"
    remote       = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `destination` (Attributes, required) (see [below for nested schema](#nestedatt--destination))
* `filename` (String, required) - File name for the generated text configuration file
* `only_traffic` (Boolean, required) - If true, restrict to traffic only configurations
* `protocol` (String, required) - File Source/File Destination protocol. http and https only applicable for fileSource
* `remote` (Boolean, required) - If true, upload the text configuration file to remote host.If false, save file to local storage

<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Required:

* `hostname` (String) - server address
* `path` (String) - configuration text file path on server
Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

