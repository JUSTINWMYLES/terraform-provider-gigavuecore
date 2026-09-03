resource "gigavuecore_group" "example" {
  description = "example-value"
  name        = "example"
  roles       = ["example"]
  tags = [{
    multi_valued  = true
    override_user = true
    tag_key       = "example"
    tag_values    = ["example"]
  }]
}
