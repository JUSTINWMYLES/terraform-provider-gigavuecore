resource "gigavuecore_elb" "example" {
  alias = "example"
  hash_fields = [{
    hash_field    = "ip"
    hash_location = "inner"
  }]
}
