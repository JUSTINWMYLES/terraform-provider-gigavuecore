action "gigavuecore_redefine_snmp_server_nofify_config" "example" {
  config {
    cluster_id    = "example"
    notify_events = ["secondflashboot"]
    notify_set    = "all"
  }
}
