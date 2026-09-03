action "gigavuecore_create_deploy" "example" {
  config {
    body = {
      name          = "example"
      platform_type = "aws"
    }
    env_id = "example"
  }
}
