resource "gigavuecore_role" "example" {
  description = "examplee"
  name        = "example"
  scope = [{
    actions   = [ "all" ]
    hierarchy = true
    type      = "example"
  }]
}
