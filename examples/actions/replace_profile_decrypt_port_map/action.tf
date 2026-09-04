action "gigavuecore_replace_profile_decrypt_port_map" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    port_maps = [{
      in_port  = 1
      out_port = 1
      rule_id  = 0
    }]
  }
}
