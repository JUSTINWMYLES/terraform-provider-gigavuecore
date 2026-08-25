action "gigavuecore_update_keystore_preference" "example" {
  config {
    auto_delete     = true
    auto_enable     = true
    auto_purge      = true
    body_cluster_id = "example"
    cluster_id      = "example"
    max_keys        = 1
    retention_time  = 1
  }
}
