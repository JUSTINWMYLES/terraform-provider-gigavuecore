resource "gigavuecore_tag" "example" {
  description  = "example-value"
  hierarchical = true
  multi_valued = true
  tag_key      = "example"
  tag_type     = "Rbac"
  tag_values   = ["example"]
}
