resource "gigavuecore_ib_pathway" "example" {
  alias        = "example"
  comment      = "example"
  min_ports_up = 1
  ports        = [ "example" ]
  traffic_path = "bypass"
}
