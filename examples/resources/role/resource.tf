resource "gigavuecore_role" "example" {
  description = "example"
  name        = "example"
  scope = [{
    actions   = [ "example" ]
    hierarchy = true
    type      = "example"
  }]
}
