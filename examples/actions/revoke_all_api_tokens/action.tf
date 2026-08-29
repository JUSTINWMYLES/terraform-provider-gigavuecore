action "gigavuecore_revoke_all_api_tokens" "example" {
  config {
    context = {
      page_no     = 0
      page_size   = 0
      sort        = [ "example" ]
      total_items = 0
    }
    fm_api_token_user_entities = [{
      authentication_type = "example"
      created_by          = "example"
      created_ts          = "example"
      expiry_time         = "example"
      expiry_ts           = "example"
      groups              = [ "example" ]
      token               = "example"
      token_id            = "example"
      token_name          = "example"
      usage_count         = "example"
      username            = "example"
    }]
  }
}
