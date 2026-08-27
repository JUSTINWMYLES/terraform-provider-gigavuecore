resource "gigavuecore_port_filter" "example" {
  cluster_id = "example"
  port       = "example"
  rules = {
    drop_rules = null
    pass_rules = null
  }
}
