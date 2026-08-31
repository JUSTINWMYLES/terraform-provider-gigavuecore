---
page_title: "gigavuecore_upload_flow_ops_report Action - gigavuecore"
subcategory: ""
description: |-
  Upload a FlowOps Report
---

# gigavuecore_upload_flow_ops_report Action

Upload a FlowOps Report

## Example Usage

```terraform
action "gigavuecore_upload_flow_ops_report" "example" {
  config {
    alias                = "example"
    caller_id_pattern    = "example"
    cluster_id           = "example"
    flow_ops_report_type = "inlineSsl"
    upload_destination = {
      hostname = "example"
      password = "example"
      path     = "example"
      protocol = "scp"
      username = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `caller_id_pattern` (String, optional) - valid only if 'flowOpsReportType' is 'flowSip'.
* `cluster_id` (String, required) - Target Cluster ID
* `flow_ops_report_type` (String, required)
* `upload_destination` (Attributes, required) - Remote file source or destination (see [below for nested schema](#nestedatt--upload_destination))

<a id="nestedatt--upload_destination"></a>
### Nested Schema for `upload_destination`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

