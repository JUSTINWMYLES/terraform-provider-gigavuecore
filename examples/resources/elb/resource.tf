resource "gigavuecore_elb" "example" {
  alias = "example"
  hash_fields = [{
    hash_field    = "example"
    hash_location = "example"
  }]
}
