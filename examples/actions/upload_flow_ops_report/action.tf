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
