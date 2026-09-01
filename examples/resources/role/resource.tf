resource "gigavuecore_role" "example" {
  description = "example-value"
  name        = "example"
  scope = [{
    actions   = ["all"]
    hierarchy = true
    type      = "example"
  }]
}
