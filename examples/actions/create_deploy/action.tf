action "gigavuecore_create_deploy" "example" {
  config {
    body = {
      name          = "example"
      platform_type = "example"
    }
    env_id = "example"
  }
}
