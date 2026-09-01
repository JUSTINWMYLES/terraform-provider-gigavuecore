action "gigavuecore_service_state_change" "example" {
  config {
    hostname = "example"
    service  = "application"
    state    = "up"
  }
}
