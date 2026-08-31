action "gigavuecore_update_email_notif_config_spec" "example" {
  config {
    cluster_id    = "example"
    notify_events = [ "secondflashboot" ]
    notify_set    = "all"
  }
}
