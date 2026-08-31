resource "gigavuecore_tag" "example" {
  description  = "examplee"
  hierarchical = true
  multi_valued = true
  tag_key      = "example"
  tag_type     = "Rbac"
  tag_values   = [ "example" ]
}
