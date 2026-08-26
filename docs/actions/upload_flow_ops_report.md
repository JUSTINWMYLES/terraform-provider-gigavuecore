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
    flow_ops_report_type = "example"
    upload_destination   = "example"
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
* `upload_destination` (Dynamic, required) - Remote file source or destination


