resource "gigavuecore_hsm_group" "example" {
  alias              = "example"
  comment            = "example"
  hsms               = [ "example" ]
  operational_status = "unknown"
  type               = "ncipher"
}
