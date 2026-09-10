action "gigavuecore_update_deploy" "example" {
  config {
    body = {
      name          = "example"
      platform_type = "aws"
    }
    env_id   = "example"
    unify_id = "example"
  }
}
