action "gigavuecore_add_key_map_to_hsm_group" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    hsm_key_maps = [{
      address    = "example"
      cluster_id = "example"
      key_name   = "example"
      key_token  = "example"
      port       = 0
      rfs_match  = "example"
      rule_id    = "example"
    }]
  }
}
