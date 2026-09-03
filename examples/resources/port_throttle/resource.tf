resource "gigavuecore_port_throttle" "example" {
  alias = "example"
  ports_throttles = [{
    port  = "example"
    type  = "pps"
    value = 33
  }]
}
