---
page_title: "gigavuecore_restore_node_system_config_text_file Action - gigavuecore"
subcategory: ""
description: |-
  Apply the text configuration to the running system from a remote host or local storage
---

# gigavuecore_restore_node_system_config_text_file Action

Apply the text configuration to the running system from a remote host or local storage

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_restore_node_system_config_text_file" "example" {
  config {
    clear_config = true
    cluster_id   = "example"
    destination = {
      hostname = "example"
      password = "example"
      path     = "example"
      username = "example"
    }
    fail_continue = true
    filename      = "example"
    protocol      = "scp"
    remote        = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `clear_config` (Boolean, optional) - If true, clear the existing traffic configuration before applying the text configuration file.
* `cluster_id` (String, required) - Target Cluster ID
* `destination` (Attributes, required) (see [below for nested schema](#nestedatt--destination))
* `fail_continue` (Boolean, required) - If true, while applying commands, continue execution even if one of them fails
* `filename` (String, required) - File name of the text configuration file to be applied on system
* `protocol` (String, required) - File Source/File Destination protocol. http and https only applicable for fileSource
* `remote` (Boolean, required) - If true, fetch the text configuration file from remote host.If false, fetch file from local storage

<a id="nestedatt--destination"></a>
### Nested Schema for `destination`

Required:

* `hostname` (String) - server address
* `path` (String) - configuration text file path on server

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

