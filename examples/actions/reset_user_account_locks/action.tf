action "gigavuecore_reset_user_account_locks" "example" {
  config {
    clear_history = true
    cluster_id = "example"
    unlock = true
    username = "example"
  }
}
