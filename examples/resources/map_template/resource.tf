resource "gigavuecore_map_template" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  roles = {
    editors = [ "example" ]
    owners  = [ "example" ]
    viewers = [ "example" ]
  }
  rules = {
    drop_rules = null
    pass_rules = null
  }
}
