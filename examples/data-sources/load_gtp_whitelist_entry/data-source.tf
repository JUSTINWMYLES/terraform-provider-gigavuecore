data "gigavuecore_load_gtp_whitelist_entry" "example" {
  alias      = "example"
  cluster_id = "example"
  imsi       = "12345678901234"
  ran        = "123.45.0xabcd"
}
